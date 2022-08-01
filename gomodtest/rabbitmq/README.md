# 多个消费者订阅同一个 queue
和 nsq 多个消费者订阅同一个 channel 一样，会负载均衡的推送给消费者

# 生产者和消费者的 QueueDeclare() 传递的参数不一样会怎样，比如 durable 不同
QueueDeclare() 会先检查要创建的 queue 是否存在，如果不存在，则根据提供的参数创建，如果已存在，
则检查提供的参数是否一致，不一致会报错

比如现在已经存在一个名为 hello 的 queue，它的 durable 为 false，现在我在生产者处调用：
```go
q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
```
传递的 durable 为 true

运行后的结果为：
```shell
Failed to declare a queue: Exception (406) Reason: "PRECONDITION_FAILED - inequivalent arg 'durable' for queue 'hello' in vhost '/': received 'true' but current is 'false'"
panic: Failed to declare a queue: Exception (406) Reason: "PRECONDITION_FAILED - inequivalent arg 'durable' for queue 'hello' in vhost '/': received 'true' but current is 'false'"
```

# 队列持久化和消息持久化为什么要区分开？既然消息是属于队列的，那么队列声明为持久化，消息应该不也默认为持久化吗？


# 怎么解决重复消费


# 怎么保证可靠性
为了确保消息永不丢失，RabbitMQ 支持 [消息确认](https://www.rabbitmq.com/confirms.html)。
消费者发送回一个确认（acknowledgement），以告知 RabbitMQ 已经接收，处理了特定的消息，
并且 RabbitMQ 可以自由删除它。

如果使用者在不发送确认的情况下死亡（其通道已关闭，连接已关闭或 TCP 连接丢失），RabbitMQ 将了解
消息未完全处理，并将对其重新排队。如果同时有其他消费者在线，它将很快将其重新分发给另一个消费者。
这样，您可以确保即使工人偶尔死亡也不会丢失任何消息。

# 生产者推送消息时，没有订阅的消费者，消息会丢失吗？

# Consume() 里面有部分参数和 QueueDeclare() 相同，比如都有 exclusive 这个字段，二者有什么区别吗？