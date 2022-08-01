package main

import (
	"bufio"
	"bytes"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	. "rabbitmq"
)

func main() {
	var exchangeName = "logs_direct"

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
		exchangeName,        // name 交换器名
		ExchangeKind_Direct, // type 交换器类型
		false,               // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	Assert(err, "Failed to declare a exchange")

	log.Printf("Enter q or Q to exit.")
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
		before, after, found := bytes.Cut(body, []byte(":"))
		if !found {
			log.Printf("usage: <routing_key>:<message>")
			continue
		}

		// 4.将消息发布到 exchange
		err = ch.Publish(
			exchangeName,   // exchange
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
