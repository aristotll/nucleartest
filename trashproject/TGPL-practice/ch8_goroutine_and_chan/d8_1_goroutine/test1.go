package main

import (
	"fmt"
	"time"
)

func g1() {
	for i := 0; i < 100; i++ {
		fmt.Println("g1: ", i)
		time.Sleep(time.Second / 2)
	}
}

func g2() {
	for i := 0; i < 100; i++ {
		fmt.Println("g2: ", i)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	go g1()
	go g2()

	for i := 0; i < 100; i++ {
		fmt.Println("main: ", i)
		time.Sleep(time.Second / 2)
	}
}
