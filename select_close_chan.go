package main

import (
	"fmt"
)

var done = make(chan struct{})

func fn(x, y chan int) {
	go func() {
		for {
			select {
			case _, ok := <- x:
				if !ok {x = nil}
				fmt.Println("x")
			case _, ok := <- y:
				if !ok {y = nil}
				fmt.Println("y")
			default:
				fmt.Println("default")
				done <- struct{}{}		
			}
		}
	}()	
}

func main() {
	x := make(chan int)
	y := make(chan int)
	fn(x, y)
	x <- 1
	y <- 1
	close(x)
	close(y)
	<-done
}
