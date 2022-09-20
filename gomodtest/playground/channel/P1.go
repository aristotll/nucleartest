package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Producer(c chan int) {
	defer wg.Done()
	rand.Seed(time.Now().Unix())
	i := rand.Intn(100)
	fmt.Println("生产的随机数：", i)
	c <- i
}

func Consumer(c chan int) {
	defer wg.Done()
	var v = <- c
	var sum = 0
	for v > 0 {
		sum += v % 10
		fmt.Printf("取出一位： %d ", v % 10)
		v /= 10
	}
	fmt.Println(sum)
}

func main() {
	ch := make(chan int)
	// wg.Add(1)
	go Producer(ch)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go Consumer(ch)
	}
	wg.Wait()
}
