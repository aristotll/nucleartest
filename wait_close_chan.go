package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int)

	wg.Add(2)
	go func() {
		defer wg.Done()
		<-c // 等待 c 中有数据
		fmt.Println("ok")
	}()

	go func() {
		defer wg.Done()
		<-time.After(time.Second * 2)
		close(c) // 2 秒后关闭 c，看 <-c 处会发生什么
	}()
	wg.Wait()
	
}
