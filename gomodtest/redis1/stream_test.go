package redis1

import (
	"context"
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

var stream = NewStream(Conn(), &Config{ShowCmd: true})

func TestXADD(t *testing.T) {
	m := map[string]string{
		"name": "zhang3",
		"age":  "18",
	}
	ctx := context.Background()
	result, err := stream.ProduceMessage(ctx, "mq", m)
	if err != nil {
		panic(err)
	}
	log.Println(result)
}

func TestXREAD(t *testing.T) {
	ctx := context.Background()
	// stream 必须以 数字（适用于非阻塞状态，数字是开始读取的消息序号）
	// 或者 $ （适用于阻塞状态，获取最新的消息 ID）结尾
	xread, err := stream.ConsumeMessage(ctx,
		[]string{"mq"},
		10,
		&BlockArgs{NoBlock: &struct{ readStartId string }{readStartId: "0"}})
	if err != nil {
		panic(err)
	}
	for _, stream := range xread {
		log.Printf("%+v\n", stream)
	}

	xread, err = stream.ConsumeMessage(ctx,
		[]string{"mq"},
		10,
		&BlockArgs{Block: &struct{ maxWaitTime time.Duration }{maxWaitTime: time.Second * 5}})
	if err != nil {
		panic(err)
	}
	for _, stream := range xread {
		log.Printf("%+v\n", stream)
	}
}

func TestXGROUP(t *testing.T) {
	groupName := "g1"
	stream_ := "mq"
	err := stream.CreatGroup(context.Background(), groupName, stream_, "0")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v create ok\n", groupName)
}

func TestXINFO_GROUP(t *testing.T) {
	stream_ := "mq"
	info, err := stream.GroupInfo(context.Background(), stream_)
	if err != nil {
		t.Log("err: ", err.Error())
	}
	t.Logf("%+v\n", info)
}

func Test_XREADGROUP(t *testing.T) {
	groupName := "mq"
	consumerName := "zhang3"
	count := int64(1)
	streams := []string{"mq"}
	res, err := stream.GroupConsume(context.Background(), groupName, consumerName,
		count, streams, 0, false)
	if err != nil {

	}
}
