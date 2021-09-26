package main

import (
	"sync"
	"fmt"
)

func main() {
	var mu sync.Mutex
	mu.Unlock()
	fmt.Println("done")
}
