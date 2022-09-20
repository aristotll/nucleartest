package main //golang 唯一的神

import (
	"fmt"
	"time"
)

var oneDone = make(chan int)
var twoDone = make(chan int)

func first(n int) {
	for i := 0; i < n; i++ {
		<-twoDone
		fmt.Println("foo")
		oneDone <- 1
	}
}
func seccond(n int) {
	for i := 0; i < n; i++ {
		<-oneDone
		fmt.Println("bar")
		twoDone <- 1
	}
}

func main() {
	n := 3 //n为打印次数
	go first(n)
	go seccond(n)
	twoDone <- 1            //2个线程启动后才开始输出
	time.Sleep(time.Second) //防止主线程退出后，子线程没运行完毕
}
