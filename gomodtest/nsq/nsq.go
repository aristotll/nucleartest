package nsq

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

var producer *nsq.Producer

const (
	topic             = "room1"
	nsqdAddr          = "127.0.0.1:4150"
	nsqLookupHttpAddr = "127.0.0.1:4161"
	nsqLookupTcpAddr  = "127.0.0.1:4160"
)

func init() {
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		panic(err)
	}
	producer = p
}

func production() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "q" || text == "Q" {
			fmt.Println("bye~")
			break
		}
		if err := producer.Publish(topic, []byte(text)); err != nil {
			log.Println("publish error: ", err)
			break
		}
	}
}

type MyHandler struct {
	Name string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("[%v] recv from %v, msg: %v\n", m.Name, msg.NSQDAddress, string(msg.Body))
	return nil
}

var channel1 = "userId:1"
var channel2 = "userId:2"

// topic: keepalive | channel: chan1 | ConnectToNSQD()
func consumption1() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel1, config)
	if err != nil {
		log.Println(err)
		return
	}
	c := &MyHandler{"consumer1"}
	consumer.AddHandler(c)
	if err := consumer.ConnectToNSQD(nsqdAddr); err != nil {
		log.Println(err)
		return
	}
	select {}
}

// topic: keepalive | channel: chan2 | ConnectToNSQD()
func consumption2() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel2, config)
	if err != nil {
		log.Println(err)
		return
	}
	c := &MyHandler{"consumer2"}
	consumer.AddHandler(c)
	// 这个函数是通过 http 来通信的，所以需要传入 nsqlookup http 的监听地址
	if err := consumer.ConnectToNSQLookupd(nsqLookupHttpAddr); err != nil {
		log.Println(err)
		return
	}
	select {}
}

func main() {
	production()
}
