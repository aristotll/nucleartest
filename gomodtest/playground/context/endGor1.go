package main

import (
	"fmt"
	"time"
)

type obj int


func (o *obj) handle(c chan bool) {
	for {
		fmt.Printf("func::")
		select {
		case <-c:
			fmt.Println("结束支付")
			return
		default:
			fmt.Println("等待用户支付...")
			<-time.Tick(time.Second)
		}
	}
}

func main() {
	var o obj
	c := make(chan bool)
	go o.handle(c)
	// 睡眠 main 线程，保证 handle 能被执行
	time.Sleep(time.Second * 3)
	c <- true
	fmt.Println("main")
	// 继续睡眠 main 线程，保证 c <- true 执行
	time.Sleep(time.Second * 3)
}
