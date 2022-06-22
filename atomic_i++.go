package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	_ "sync/atomic"
	"time"
)

var a, b int64 = 1, 2

func main() {
	var wg sync.WaitGroup
	var count = 50
	wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			atomic.CompareAndSwapInt64(&n, n, n+1)
			a = 3
			b = a
		}()
		fmt.Println(atomic.LoadInt64(&n))

		go func() {
			defer wg.Done()
			fmt.Printf("a = %d, b = %d \n", a, b)
		}()

		time.Sleep(time.Millisecond * 500)
	}

	wg.Wait()
}
