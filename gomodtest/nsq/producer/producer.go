package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	nsq_ "nsqtest"
	"os"
)

func Produce(producer *nsq.Producer, topic string, body []byte) error {
	return producer.Publish(topic, body)
}

// 需要满足以下正则表达式
// ^[\.a-zA-Z0-9_-]+(#ephemeral)?$
var topic = flag.String("t", "", "topic")

func main() {
	flag.Parse()
	if *topic == "" {
		fmt.Println("usage: -t [topic]")
		return
	}
	fmt.Println("topic: ", *topic)

	config := nsq.NewConfig()
	p, err := nsq.NewProducer(nsq_.NsqdAddr, config)
	if err != nil {
		panic(err)
	}

	fmt.Println("if you want to quit, please input q or Q")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "q" || text == "Q" {
			fmt.Println("bye~")
			break
		}
		if err := Produce(p, *topic, []byte(text)); err != nil {
			log.Println("produce error: ", err)
			break
		}
	}
}
