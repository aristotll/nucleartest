package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var flag = true
var count int64 = 0
var wg sync.WaitGroup

func handle() {
	wg.Done()
	if flag {
		fmt.Println("等待用户支付...", count)
		<-time.Tick(time.Second / 200)
		// 原子 + 1
		atomic.AddInt64(&count, 1)
	}else {
		fmt.Println("取消支付")
	}
}

func main() {
	for {
		wg.Add(1)
		go handle()
		// 取原子值
		v := atomic.LoadInt64(&count)
		if v == 20 {
			flag = false
			break
		}
		<-time.Tick(time.Second)
	}
	wg.Wait()

}