package main

import "time"

func main() {
	c := make(chan int)

	go func() {
		c <- 1
		close(c)
	}()

	time.Sleep(time.Second * 2)

	<-c
}
