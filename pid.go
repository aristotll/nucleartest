package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Printf("[goroutine1] os.Getpid(): %v\n", os.Getpid())
	}()
	go func() {
		defer wg.Done()
		fmt.Printf("[goroutine2] os.Getpid(): %v\n", os.Getpid())
	}()
	fmt.Printf("[main] os.Getpid(): %v\n", os.Getpid())
	wg.Wait()

	// Output:
	// [main] os.Getpid(): 23053
	// [goroutine2] os.Getpid(): 23053
	// [goroutine1] os.Getpid(): 23053
}
