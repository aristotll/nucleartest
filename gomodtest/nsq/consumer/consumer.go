package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	nsq_ "nsqtest"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type MyHandler struct {
	Name string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("[%v] recv from %v, msg: %v\n", m.Name, msg.NSQDAddress, string(msg.Body))
	return nil
}

var (
	// 这两个的值都有要求，需要满足以下正则表达式
	// ^[\.a-zA-Z0-9_-]+(#ephemeral)?$
	topic   = flag.String("t", "", "topic")
	channel = flag.String("c", "", "channel")
)

func main() {
	flag.Parse()
	if *topic == "" || *channel == "" {
		fmt.Printf("usage: -t [topic] -c [channel]\n")
		return
	}
	fmt.Printf("topic: %v, channel: %v\n", *topic, *channel)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(*topic, *channel, config)
	if err != nil {
		log.Println(err)
		return
	}
	c := &MyHandler{"consumer1"}
	consumer.AddHandler(c)
	// 这个函数是通过 http 来通信的，所以需要传入 nsqlookup http 的监听地址
	if err := consumer.ConnectToNSQLookupd(nsq_.NsqLookupHttpAddr); err != nil {
		log.Println(err)
		return
	}

	signch := make(chan os.Signal, 1)
	signal.Notify(signch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	select {
	case sig := <-signch:
		log.Println(sig.String())
		return
	}
}
