package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for {
			select {
			case ch <- 1:
				fmt.Println("send ok")
			default:
				fmt.Println("default")
				time.Sleep(time.Second * 2)
			}
		}
	}()
	time.Sleep(time.Second * 10)
	for i := 0; i < 5; i++ {
		<-ch
		fmt.Println("get ok")
		time.Sleep(time.Second * 10)
	}

	wg.Wait()
}
