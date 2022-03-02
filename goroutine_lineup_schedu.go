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
		defer wg.Done()
		fmt.Println(n)
	}(1)

	go func(n int) {
		defer wg.Done()
		fmt.Println(n)
	}(2)

	go func(n int) {
		defer wg.Done()
		fmt.Println(n)
	}(3)

	wg.Wait()
}
