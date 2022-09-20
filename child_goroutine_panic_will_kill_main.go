package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	go func() {
		fmt.Println("panic goroutine")
		panic("!!!")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("main over after child goroutine panic")
}
