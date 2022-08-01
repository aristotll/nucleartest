package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	dictionary := make(map[int]string)
	for i := 0; i < 10000000; i++ {
		dictionary[i] = "hello"
	}
	duration := time.Since(start)
	fmt.Println("time used:", duration)
}
