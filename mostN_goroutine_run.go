package main

import (
	"fmt"
	"sync"
	"time"
)

// 最多只能有 N 个 goroutine 同时运行
func main() {
	var (
		wg          sync.WaitGroup
		num         = 100
		mostRunning = 10
		ch          = make(chan struct{}, mostRunning)
	)
	wg.Add(num)

	for i := 0; i < num; i++ {
		i := i
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			ch <- struct{}{}
			fmt.Printf("%d is running...\n", i)
			time.Sleep(time.Second * 3)
		}()
	}
	wg.Wait()
}
