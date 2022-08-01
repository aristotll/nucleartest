package redis1

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

func setNX(rdb *redis.Client, key, val, serverName string) {
	bc := rdb.SetNX(context.Background(), key, val, time.Second*5)
	ok, _ := bc.Result()
	done := make(chan struct{})

	go func() {
		if ok {
			select {
			case <-time.After(time.Second * 5):
				del(rdb, key)
				fmt.Printf("%s 超时，被强制释放锁 \n", serverName)
			case <-done:
				del(rdb, key)
				fmt.Printf("%s 正常释放锁 \n", serverName)
			}
		}
	}()

	if ok {
		fmt.Printf("%s 抢到锁了 \n", serverName)
		fmt.Println("处理业务中...")
		time.Sleep(time.Second * 3)
		fmt.Println("处理完成")
		done <- struct{}{} // 通知处理完成
	}
}

func del(rdb *redis.Client, key string) {
	_, err := rdb.Del(context.Background(), "testgeneric").Result()
	if err != nil {
		log.Fatalln(err)
	}
}
