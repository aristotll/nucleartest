// map 能否并发读？

package main

import (
	"fmt"
	"time"
)

// 结论：map 可以并发读
func main() {
	m := make(map[int]int)
	m[1] = 1
	count := 100
	for i := 0; i < count; i++ {
		go func (i int) {
			fmt.Printf("[%d] m: %v\n", i, m[1])
		}(i)
	}

	time.Sleep(time.Second * 10)
}