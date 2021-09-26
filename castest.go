package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int64 = 10
	b := atomic.CompareAndSwapInt64(&i, 10, 50)
	fmt.Printf("b: %v\n", b)
}