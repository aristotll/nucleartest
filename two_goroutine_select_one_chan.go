package main

import (
	"fmt"
	//"sync"
)

func main() {
	c := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(2)

	go func() {
		//defer wg.Done()
		for {
			select {
			case <-c:
				fmt.Println("goroutine1")
			}
		}
	}()

	go func() {
		//defer wg.Done()
		for {
			select {
			case <-c:
				fmt.Println("goroutine2")
			}
		}
	}()

	for i := 0; i < 100; i++ {
		c <- 1
	}

	//wg.Wait()
}
