package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 100)

	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println(v)
			}
		}
	}()

	time.Sleep(time.Second * 3)

	go func() {
		ch1 <- 100
	}()

	time.Sleep(time.Second * 10)
}
