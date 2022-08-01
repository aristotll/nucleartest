package main

import (
	"bufio"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	. "rabbitmq"
)

func main() {
	// 1. 尝试连接 RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数 API 都是用过该通道操作的。
	ch, err := conn.Channel()
	Assert(err, "Failed to open a channel")
	defer ch.Close()

	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable 是否为持久化队列
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	Assert(err, "Failed to declare a queue")

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
		// 4.将消息发布到声明的队列
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			// 如果为 true, 会根据 exchange 类型和 routkey 规则，如果无法找到符合条件的队列
			// 那么会把发送的消息返回给发送者
			false, // mandatory
			// 如果为 true, 当 exchange 发送消息到队列后发现队列上没有绑定消费者，
			// 则会把消息发还给发送者
			false, // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
				// 是否持久化消息，瞬态（0 或 1）或持久（2）
				DeliveryMode: amqp.Transient,
			})
		Assert(err, "Failed to publish a message")
	}
}
