package main

import "fmt"

func rNum() <-chan int {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		c <- i
	}
	// 必须关闭，不然接收端会死锁
	defer close(c)
	return c
}

func main() {
	num := rNum()
	for v := range num {
		fmt.Println(v)
	}
}
