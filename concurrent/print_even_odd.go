package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	odd := 1
	even := 2
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			<-ch1
			fmt.Println(odd)
			time.Sleep(time.Second)
			odd += 2
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			<-ch2
			fmt.Println(even)
			time.Sleep(time.Second)
			even += 2
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()
}
