package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var num int64

func main() {
	var (
		wg sync.WaitGroup
		n  int64 = 1000
	)
	wg.Add(int(n))

	for i := 0; int64(i) < n; i++ {
		go func() {
			defer wg.Done()
			for !atomic.CompareAndSwapInt64(&num, num, num+1) {}
			fmt.Println(num)
		}()
	}
	
	wg.Wait()
	fmt.Println("result: ", num)
}
