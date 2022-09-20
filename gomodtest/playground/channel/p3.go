package main

import (
	"fmt"
	"sync"
	"time"
)

// 顺序打印字母

var wg sync.WaitGroup

func pr1(c chan<- string) {
	fmt.Println("我先打印")
	c <- "Hello, World! dasdsadas 带房间辣豆腐 就是个ID结构 🏷👨‍🎓🐀💑"
	defer wg.Done()
}

func pr2(c <-chan string) {
	v := <-c
	for _, val := range v {
		fmt.Printf("%c ", val)
		time.Sleep(time.Second / 5)
	}
	defer wg.Done()
}

func main() {
	c := make(chan string)

	wg.Add(2)

	go pr1(c)
	go pr2(c)

	wg.Wait()

}
