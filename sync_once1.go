package main

import (
	"sync"
	"fmt"
)

func main() {
	once := sync.Once{}
	for i := 0; i < 100; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
