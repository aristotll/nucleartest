package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			ch <- 1
			time.Sleep(time.Second)
		}
		close(ch) // 一定要 close，不然 range 会死锁
	}()

	for v := range ch {
		fmt.Println(v)
	}

	wg.Wait()
}
