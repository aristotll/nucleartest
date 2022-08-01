package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math/rand"
	. "rabbitmq"
	"sync"
	"time"
)

func main() {
	var (
		queueName    = "order_q"
		dlxQueueName = "dlx_order_q"
		wg           sync.WaitGroup
	)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Assert(err, "Failed to open a channel")
	defer ch.Close()

	wg.Add(2)

	rand.Seed(time.Now().Unix())
	// 消费正常队列
	go func() {
		defer wg.Done()
		msgs, err := ch.Consume(
			queueName,
			"",
			false,	// auto-ack
			false,
			false,
			false,
			nil,
		)
		Assert(err, "Failed to consume queue")
		for msg := range msgs {
			log.Printf("收到订单：%s\n", msg.Body)
			//  RabbitMQ 不会为未确认的消息设置过期时间，它判断此消息是否需要重新投递给消费者的
			// 唯一依据是消费该消息的消费者连接是否己经断开，这么设计的原因是 RabbitMQ 允许消费者
			// 消费一条消息的时间可以很久很久。
			t := time.Duration(rand.Intn(5) + 10) * time.Second
			log.Printf("处理所需时间：%v\n", t)
			time.Sleep(t)
			msg.Ack(false)
			// 第二个参数：requeue：代表是否将这条消息重新放入队列，如果重新放回，这可以被其他消费者
			// 继续消费，否则这条消息会丢失
			// msg.Nack(false, true)	
		}
	}()

	// 消费死信队列
	go func() {
		defer wg.Done()
		msgs, err := ch.Consume(
			dlxQueueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		Assert(err, "Failed to consume queue")
		for msg := range msgs {
			log.Printf("订单过期：%v，当前时间：%v\n",
				string(msg.Body), time.Now().Format("2006-01-02 15:04:05"))
		}
	}()

	wg.Wait()
}
