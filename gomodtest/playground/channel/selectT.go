package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("commencing countdown. ")
	//
	tick := time.Tick(time.Second * 1)
	for countdown := 0; countdown < 10; countdown++ {
		fmt.Println(countdown)
		<-tick
	}

	// time.After()
}
