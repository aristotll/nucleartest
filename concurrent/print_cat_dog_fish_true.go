package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// cat, dog, fish
	ch1, ch2, ch3 := newChan(), newChan(), newChan()
	var count int64
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&count) < 100 {
			<-ch1
			fmt.Println("cat")
			atomic.AddInt64(&count, 1)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&count) < 100 {
			<-ch2
			fmt.Println("dog")
			atomic.AddInt64(&count, 1)
			ch3 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for atomic.LoadInt64(&count) < 100 {
			<-ch3
			fmt.Println("fish")
			atomic.AddInt64(&count, 1)
			if atomic.LoadInt64(&count) < 100 {
				ch1 <- struct{}{}
			}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()
}

func newChan() chan struct{} {
	return make(chan struct{})
}
