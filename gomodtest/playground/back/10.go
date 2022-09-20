package main

import (
	"fmt"
	"time"
)

var m = make(map[int]bool)

func printOnce(key int) {
	if _, ok := m[key]; !ok {
		fmt.Println(key)
	}
	m[key] = true
}


func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	<-time.Tick(time.Second)
}
