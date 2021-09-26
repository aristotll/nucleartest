package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg sync.WaitGroup
		ch1 = make(chan struct{})
		ch2 = make(chan struct{})
		ch3 = make(chan struct{})
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i += 3 {
			<-ch1
			fmt.Println("goroutine1: ", i + 1)
			if i+3 < 1000 {
				ch2 <- struct{}{}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 1000; i += 3 {
			<-ch2
			fmt.Println("goroutine2: ", i + 1)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i < 1000; i += 3 {
			<-ch3
			fmt.Println("goroutine3: ", i + 1)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}
	wg.Wait()
}
