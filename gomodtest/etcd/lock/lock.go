package main

import (
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"sync"
	"time"
)

const lockerPfx = "/locker/"

func fn(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Second * 10,
	})
	if err != nil {
		log.Println(err)
		return
	}

	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	locker := concurrency.NewLocker(session, lockerPfx)

	log.Printf("[%v] 尝试获取锁 %v \n", endpoint, lockerPfx)
	locker.Lock()
	log.Printf("[%v] 获得锁 %v \n", endpoint, lockerPfx)

	log.Printf("[%v] 处理业务中... \n", endpoint)
	time.Sleep(time.Second * 5)
	log.Printf("[%v] 处理业务完成... \n", endpoint)

	locker.Unlock()
	log.Printf("[%v] 释放锁 %v \n", endpoint, lockerPfx)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go fn("127.0.0.1:2379", &wg)
	go fn("127.0.0.1:22379", &wg)
	go fn("127.0.0.1:32379", &wg)

	wg.Wait()
}
