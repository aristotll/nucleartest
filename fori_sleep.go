package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		fmt.Print(1)
		time.Sleep(time.Second)
	}
}
