package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Unlock()
	fmt.Println("done")
}
