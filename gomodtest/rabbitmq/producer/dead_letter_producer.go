package main

import (
	"bufio"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	. "rabbitmq"
	"sync/atomic"
	"time"
)

func main() {
	var (
		exchangeName    = "order"
		dlxExchangeName = "dead_letter_exchange"
		routingKey      = "order_routing"
		dlxRoutingKey   = "dlx_routing"
		queueName       = "order_q"
		dlxQueueName    = "dlx_order_q"
	)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Assert(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个正常的 direct 类型的交换器
	err = ch.ExchangeDeclare(
		exchangeName,        // Exchange names
		amqp.ExchangeDirect, //"direct", "fanout", "topic" and "headers"
		true,
		false, // Durable and Non-Auto-Deleted exchanges 会一直保留
		false,
		false,
		nil,
	)
	Assert(err, "ch.ExchangeDeclare err: ")

	// 声明一个用于死信消息的 direct 类型的交换器
	err = ch.ExchangeDeclare(
		dlxExchangeName,     // Exchange names
		amqp.ExchangeDirect, //"direct", "fanout", "topic" and "headers"
		true,
		false, // Durable and Non-Auto-Deleted exchanges 会一直保留
		false,
		false,
		nil,
	)
	Assert(err, "ch.ExchangeDeclare err: ")

	// 声明一个正常队列
	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		amqp.Table{
			"x-message-ttl":             5000,            // 消息过期时间,毫秒
			"x-dead-letter-exchange":    dlxExchangeName, // 指定死信交换机
			"x-dead-letter-routing-key": dlxRoutingKey,   // 指定死信 routing-key
		},
	)
	Assert(err, "ch.QueueDeclare err: ")

	// 将正常队列和正常交换机绑定
	err = ch.QueueBind(
		queue.Name,
		routingKey,
		exchangeName, // 队列绑定到正常交换机
		false,
		nil,
	)
	Assert(err, "ch.QueueBind err: ")

	// 声明一个队列用于处理死信消息
	dlxQueue, err := ch.QueueDeclare(
		dlxQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	Assert(err, "ch.QueueDeclare err: ")

	// 将死信交换机和死信队列绑定
	err = ch.QueueBind(
		dlxQueue.Name,
		dlxRoutingKey,
		dlxExchangeName, // 队列绑定到正常交换机
		false,
		nil,
	)
	Assert(err, "ch.QueueBind err: ")

	log.Printf("Enter any content to generate an order.\nEnter q or Q to exit.")
	var orderId atomic.Value
	orderId.Store(int64(1))

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

		id := orderId.Load().(int64)
		msg := fmt.Sprintf("订单编号：%v, 订单生成时间：%v",
			id, time.Now().Format("2006-01-02 15:04:05"))
		orderId.CompareAndSwap(id, id+1)

		// 发送消息到正常交换机
		if err = ch.Publish(
			exchangeName,
			routingKey,
			false,
			false,
			amqp.Publishing{
				ContentType:  "text/plain",
				Body:         []byte(msg),
				DeliveryMode: amqp.Transient,
				Priority:     0,
			},
		); err != nil {
			fmt.Println("ch.Publish err: ", err)
			return
		}
	}
}
