package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type ExchangeKind = string

const (
	ExchangeKind_Direct  ExchangeKind = "direct"
	ExchangeKind_Topic   ExchangeKind = "topic"
	ExchangeKind_Headers ExchangeKind = "headers"
	ExchangeKind_Fanout  ExchangeKind = "fanout"
)

func Assert(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetChannel() (*amqp.Channel, error) {
	// 1. 尝试连接 RabbitMQ，建立连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等。
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	Assert(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 2. 接下来，我们创建一个通道，大多数 API 都是用过该通道操作的。
	return conn.Channel()
}

func CreateOrConnectQueue(ch *amqp.Channel) (amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable 是否为持久化队列
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	return q, err
}

func CreateOrConnectExchange(ch *amqp.Channel) {

}
