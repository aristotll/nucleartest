package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	. "rabbitmq"
)

// 检查输入的 exchange kind 是否合法
func checkExchangeKind(kind string) error {
	m := map[string]bool{
		ExchangeKind_Direct:  true,
		ExchangeKind_Topic:   true,
		ExchangeKind_Headers: true,
		ExchangeKind_Fanout:  true,
	}
	if exist := m[kind]; exist {
		return nil
	}
	return errors.New("not found this kind")
}

// ./producer_exchange -e logs_direct -k direct
// ./producer_exchange -e logs_topic -k topic
//    	> quick.orange.rabbit:666			=> 		*.orange.*  | *.*.rabbit
//		> lazy.orange.elephant:aaa			=>		*.orange.*
//		> quick.orange.fox:888				=>		*.orange.*
//		> lazy.brown.fox:666				=>		lazy.#
//		> lazy.pink.rabbit:xxx				=>		lazy.# | *.*.rabbit
//		> quick.brown.fox:yyy				=>
//		> orange:ccc						=>
// 		> quick.orange.male.rabbit:bbb		=>
//		> lazy.orange.male.rabbit:ooo		=>		lazy.#
//		> lazy.orange.male.rabbit.io:zzz	=>		lazy.#
func main() {
	log.Printf("usage: ./main -e [exchange name] -k [exchange kind]")
	var (
		exchangeName = flag.String("e", "", "exchange name")
		exchangeKind = flag.String("k", "", "exchange kind")
	)
	flag.Parse()

	if err := checkExchangeKind(*exchangeKind); err != nil {
		Assert(err, "check exchange kind error: ")
	}
	log.Printf("exchangeName: %v, exchangeKind: %v\n",
		*exchangeName, *exchangeKind)

	// 1. 尝试连接 RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数 API 都是用过该通道操作的。
	ch, err := conn.Channel()
	Assert(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个交换器 exchange
	err = ch.ExchangeDeclare(
		*exchangeName, // name 交换器名
		*exchangeKind, // type 交换器类型
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	Assert(err, "Failed to declare a exchange")

	log.Printf("Enter q or Q to exit.")
	log.Printf("usage: <routing_key>:<message>")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Err() != nil {
			Assert(scanner.Err(), "read stdin error")
		}
		body := scanner.Bytes()
		if string(body) == "q" || string(body) == "Q" {
			log.Printf("bye^")
			break
		}
		// 输入的格式：<路由键>:<消息实体>
		before, after, found := bytes.Cut(body, []byte(":"))
		if !found {
			log.Printf("usage: <routing_key>:<message>")
			continue
		}
		log.Printf("routingKey: %s, message: %s\n", before, after)

		// 4.将消息发布到 exchange 而不是某个单独的队列
		err = ch.Publish(
			*exchangeName,  // exchange
			string(before), // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        after,
				// 是否持久化消息，瞬态（0 或 1）或持久（2）
				DeliveryMode: amqp.Transient,
			})
		Assert(err, "Failed to publish a message")
	}
}
