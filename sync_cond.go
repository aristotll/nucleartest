package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	for i := 0; i < 10; i++ {
		go func() {
			cond.L.Lock()
			time.Sleep(10 * time.Second)
			cond.L.Unlock()
		}()
	}
	cond.Wait()
}
