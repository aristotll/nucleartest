package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

const subj = "room"

func Pub(conn *nats.Conn) error {
	if err := conn.Publish(subj, []byte("message 1")); err != nil {
		return err
	}
	log.Println("server: pub msg1 ok")

	if err := conn.Publish(subj, []byte("message 2")); err != nil {
		return err
	}
	log.Println("server: pub msg2 ok")
	return nil
}

func Sub(conn *nats.Conn) error {
	log.Println("sub start...")
	//sub, err := conn.SubscribeSync(subj)
	//if err != nil {
	//	return err
	//}
	//for {
	//	msg, err := sub.NextMsg(time.Second * 3)
	//	if err != nil {
	//		return err
	//	}
	//	log.Println("sub msg: ", msg.Reply)
	//}

	//ch := make(chan *nats.Msg, 64)
	//if _, err := conn.ChanSubscribe(subj, ch); err != nil {
	//	return err
	//}
	//msg := <-ch
	//log.Println("sub msg: ", msg.Data)

	_, err := conn.Subscribe(subj, func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})
	if err != nil {
		return err
	}

	select {}
	//return nil
}

func TestBasic(t *testing.T) {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	// 必须 sub 先启动，pub 后启动
	go func() {
		defer wg.Done()
		if err := Sub(conn); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 2)

	go func() {
		defer wg.Done()
		if err := Pub(conn); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

// =================== jetStream test ===================
// $ nats-server -js 开启 jetStream

func jetStreamPub() error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("connect error: ", err)
		return err
	}
	defer nc.Close()

	stream, err := nc.JetStream()
	if err != nil {
		log.Println("jet stream error: ", err)
		return err
	}
	// 必须要调用这个函数，否则会报 no response from stream
	stream.AddStream(&nats.StreamConfig{
		Name:     "room",
		Subjects: []string{"room.*"},
	})
	_, err = stream.Publish(subj, []byte("msg1"))
	if err != nil {
		log.Println("publish error: ", err)
		return err
	}
	log.Println("push msg1 ok")

	_, err = stream.Publish(subj, []byte("msg2"))
	if err != nil {
		return err
	}
	log.Println("push msg2 ok")

	return nil
}

func jetStreamSub() error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	stream, err := nc.JetStream()
	if err != nil {
		return err
	}

	// 不知道为什么，有时候能全部读出消息，有时候只能读取部分消息，有时候又完全读不出消息（空）
	_, err = stream.Subscribe(subj, func(msg *nats.Msg) {
		//msg.Ack() //
		log.Println(string(msg.Data))
	})
	if err != nil {
		return err
	}

	return nil
}

func TestJetStreamSubPull(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		t.Fatal(err)
	}
	defer nc.Close()

	stream, err := nc.JetStream()
	if err != nil {
		t.Fatal(err)
	}

	sub, err := stream.PullSubscribe(subj, "UNKNOWN")
	if err != nil {
		t.Fatal(err)
	}
	for {
		msgs, err := sub.Fetch(1)
		if err != nil {
			t.Log(err)
			break
		}
		for _, msg := range msgs {
			msg.Ack() // 防止重复消费
			log.Printf(string(msg.Data))
		}
	}
}

func TestJetStreamPub(t *testing.T) {
	if err := jetStreamPub(); err != nil {
		t.Fatal(err)
	}
}

func TestJetStreamSub(t *testing.T) {
	if err := jetStreamSub(); err != nil {
		t.Fatal(err)
	}
}

// =================== 无用的测试文件，无参考价值 ===================
// nats-server -js 开启 jetStream
func TestJetStream(t *testing.T) {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := jetStreamPub(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 2)

	go func() {
		defer wg.Done()
		if err := jetStreamSub(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func TestJetStreamBasicUsage(t *testing.T) {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Simple Stream Publisher
	_, err := js.Publish("ORDERS.scratch", []byte("hello"))
	if err != nil {
		panic(err)
	}

	// Simple Async Stream Publisher
	for i := 0; i < 500; i++ {
		js.PublishAsync("ORDERS.scratch", []byte("hello"))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Simple Async Ephemeral Consumer
	js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})

	// Simple Sync Durable Consumer (optional SubOpts at the end)
	//sub, err := js.SubscribeSync("ORDERS.*", nats.Durable("MONITOR"), nats.MaxDeliver(3))
	//m, err := sub.NextMsg(timeout)

	// Simple Pull Consumer
	//sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
	//msgs, err := sub.Fetch(10)

	//Unsubscribe
	//sub.Unsubscribe()

	// Drain
	//sub.Drain()
}
