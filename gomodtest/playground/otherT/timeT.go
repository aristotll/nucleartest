package main

import (
	"fmt"
	"time"
)

func after() {
	c := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c <- i

	}

	loop:
	for {
		select {
		case v := <-c:
			fmt.Println("server response: ", v)
			// 5 秒后，返回该 channel
			// 若进入该 case, 则表示 c 中的值已经空了
		case v := <-time.After(time.Second * 5):
			// 5s 后，会输出该语言，并跳出 for
			fmt.Println("server error 400: request out of time", v)
			break loop
		}
	}
}

func tick() {
	times := time.Tick(time.Second * 2)
	for i := 0; i < 10; i++ {
		v := <-times
		fmt.Println(v)
	}
}

func main() {
	after()
	// tick()
	//fmt.Println("start")
	//times := time.Tick(time.Second * 5)
	//<-times
	//fmt.Println(time.Now())

}
