package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 10,
	})
	if err != nil {
		log.Fatalln(err)
	}

	watchChan := cli.Watch(context.Background(), "", clientv3.WithPrefix())
	for {
		select {
		case resp := <-watchChan:
			if resp.Err() != nil {
				log.Println("watch error: ", resp.Err())
				break
			}
			for _, event := range resp.Events {
				switch event.Type {
				case clientv3.EventTypePut:
					// 新的 key
					if event.IsCreate() {
						log.Printf("[put event] new kv: key=%v, val=%v\n",
							string(event.Kv.Key), string(event.Kv.Value))
					} else if event.IsModify() {
						log.Printf(
							"[update event]: key=%v, val=%v\n",
							string(event.Kv.Key), string(event.Kv.Value))
					}
				case clientv3.EventTypeDelete:
					log.Printf("[delete event]: key=%v\n", string(event.Kv.Key))
				}
			}
		}
	}
}
