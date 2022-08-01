package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)	
	}()
	wg.Wait()
}
