package main

import (
	"fmt"
	"time"
)

var start time.Time
var end time.Time

func main() {
	test := func() {
		fmt.Print("123")
		time.Sleep(time.Second)
	}
	limitTime1(test, time.Second*5) //可以将这里的时间改为超过5s
	end = time.Now()
	fmt.Println("testgeneric()耗时：", end.Sub(start))
}

func limitTime1(test func(), duration time.Duration) {
	var ch = make(chan bool)
	go func() {
		go func() {
			//先开启计时
			<-time.NewTimer(duration).C
			//如果能执行到这里则说明test()函数超时
			ch <- true
		}()
		//再执行要监控的程序
		test()
		//如果能执行到这里则说明test()函数正常执行完，不超时
		ch <- true
	}()
	//这里阻塞等于模拟了监控计时
	<-ch
}
