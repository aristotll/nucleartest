package main

import (
	"flag"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	. "rabbitmq"
	"strings"
)

type routingKeys []string

func (s *routingKeys) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *routingKeys) Set(value string) error {
	ss := strings.Split(value, ",")
	*s = append(*s, ss...)
	return nil
}

// ./consumer_exchange -e logs_direct -r error
// ./consumer_exchange -e logs_direct -r error,info,warning
//
// ./consumer_exchange -e logs_topic -r "*.orange.*"
// ./consumer_exchange -e logs_topic -r "*.*.rabbit","lazy.#"
func main() {
	log.Printf(`
	usage: ./main 
	-q [queue name] 
	-e [exchange name] 
	-r [routing key, can provide more, use ',' to sep, 
		example: info,error,warning (don't have space) ]
`)
	var (
		exchangeName = flag.String("e", "", "exchange name")
		queueName    = flag.String("q", "", "queue name")
		rks          routingKeys
	)
	flag.Var(&rks, "r", "routing key")
	flag.Parse()

	log.Printf("queueName: %v, exchangeName: %v, routingKey: %v\n",
		*queueName, *exchangeName, &rks)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	Assert(err, "Failed to open a channel")
	defer ch.Close()

	// 声明一个临时队列，一旦消费者断开连接，该队列就会被删除
	q, err := ch.QueueDeclare(
		*queueName, // name 空字符串作为队列名称，表示使用随机名称
		false,      // durable
		false,      // delete when unused
		true,       // exclusive 独占队列（当前声明队列的连接关闭后即被删除）
		false,      // no-wait
		nil,        // arguments
	)
	Assert(err, "Failed to declare a queue")

	if len(rks) == 0 {
		err = ch.QueueBind(
			q.Name,
			"",
			*exchangeName,
			false,
			nil,
		)
		Assert(err, "Failed to bind a queue")
	} else {
		for _, key := range rks {
			// 将 queue 绑定到对应的 exchange，使用 exchangeName + routingKeys 进行匹配
			err = ch.QueueBind(
				q.Name,
				key,
				*exchangeName,
				false,
				nil,
			)
			Assert(err, "Failed to bind a queue")
		}
	}

	// 用于公平分发
	// 你可能已经注意到调度仍然不能完全按照我们的要求工作。例如，在一个有两个 worker 的情况下，
	// 当所有的奇数消息都是重消息而偶数消息都是轻消息时，一个 worker 将持续忙碌，而另一个
	// worker 几乎不做任何工作。嗯，RabbitMQ 对此一无所知，仍然会均匀地发送消息。
	//
	// 这是因为 RabbitMQ 只是在消息进入队列时发送消息。它不考虑消费者未确认消息的数量。
	// 只是盲目地向消费者发送信息。
	//
	// 为了避免这种情况，我们可以将预取计数设置为 1。这告诉 RabbitMQ 不要一次向一个
	// worker 发出多个消息。或者，换句话说，在处理并确认前一条消息之前，不要向 worker 发送
	// 新消息。相反，它将把它发送给下一个不忙的 worker。
	if err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		Assert(err, "ch.Qos() failed")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	Assert(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			// 如果您调用 Channel.Consume 时将 autoAck 设置为 true，那么服务器将自动确认
			// 每条消息，并且不应调用此方法。否则，您必须在成功处理此交付后调用 Delivery.Ack。
			// 参数是什么意思？
			// d.Ack(false)
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
