package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(1)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(2)

	go func(n int) {
		fmt.Println(n)
		wg.Done()
	}(3)

	wg.Wait()
}
