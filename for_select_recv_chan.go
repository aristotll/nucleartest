package main

import "fmt"

func main() {
	c := make(chan int)

	go func () {
		c <- 1
	}()

	for {
		select {
		case <-c:
			fmt.Println("ok")
		}
	}
}