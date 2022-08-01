package main

import (
	"fmt"
	//"sync"
)

func main() {
	//var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{})
	//wg.Add(2)

	go func() {
		for {
			<-ch3
			v := <-ch1
			fmt.Println("ch1: ", v)
			//ch3 <- struct{}{}
		}
	}()

	go func() {
		for {
			//<-ch3
			v := <-ch2
			fmt.Println("ch2: ", v)
			ch3 <- struct{}{}
		}
	}()

	for i := 0; i < 100; i++ {
		if i%2 == 1 {ch1 <- i}
		if i%2 == 0 {ch2 <- i}
	}

	//wg.Wait()
}
