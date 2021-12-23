package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	go func() {
		c <- 1
	}()
	go func() {
		fmt.Println(<-c)
	}()

	time.Sleep(time.Second * 5)
}
