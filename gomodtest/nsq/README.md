# 多个 consumer 监听同一个 channel
消息会负载均衡的发送给 consumer，比如在 topic XXX 下有一个 channel YYY，两个 consumer A， B 
同时订阅这个 YYY， 消息 1 会发送给 A，消息 2 会发送给 B，消息 3 会发送给 A，消息 4 会发送给 B，
以此类推，貌似是使用轮询的方式

# topic channel
一个 topic 下可以有多个 channel，当发送消息到 topic 时，topic 会将消息拷贝发送到每个 channel，
channel 和 consumer 关联，如果多个 consumer 关联同一个 channel，则消息会负载均衡的消费

总而言之，消息是从 topic -> channel 是多播的（每个 channel 都接收该 topic 的所有消息的副本），
但从 channel -> consumer（每个 consumer 接收该 channel 的部分消息）是均匀分布的。

# 如何保证消息的可靠性（不丢失）

# 消息重复消费

# 什么情况下会出现消息重复