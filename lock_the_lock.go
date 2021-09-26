package main

import (
	"sync"
	"fmt"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
	fmt.Println("done")
}
