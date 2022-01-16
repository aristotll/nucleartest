package main

import (
	"time"
	"fmt"
)

func main() {
	go func() {
		for {
			fmt.Println("loop...")
			time.Sleep(time.Second)
		}	
	}()

	fmt.Println("main run...")
	time.Sleep(time.Second * 10)
}
