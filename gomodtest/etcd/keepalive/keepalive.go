package keepalive

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func fn() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second * 10,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// 永久续约
	alive, err := cli.KeepAlive(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}
	for range alive {
		log.Println("keepAlive success")
	}
}
