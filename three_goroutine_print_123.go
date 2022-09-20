package main

import (
	"fmt"
	"sync"
	"time"
)

func print123() {
	var (
		ch1 = make(chan struct{})
		ch2 = make(chan struct{})
		ch3 = make(chan struct{})
		wg  sync.WaitGroup
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch1
			fmt.Print("1")
			time.Sleep(time.Second)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch2
			fmt.Print("2")
			time.Sleep(time.Second)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-ch3
			fmt.Println("3")
			time.Sleep(time.Second)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	wg.Wait()
}

func main() {
	print123()
}
