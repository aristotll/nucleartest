package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
	fmt.Println("done")
}
