package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i + 1
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
