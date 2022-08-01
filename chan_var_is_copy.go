package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)
	cc := c
	ccc := c

	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c, <-cc, <-ccc)
}
