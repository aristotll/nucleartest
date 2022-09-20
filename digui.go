package main

import (
	"fmt"
)

func digui(i int) {
	ii := 123
	if i < 0 {
		return
	}
	ii++
	fmt.Printf("i: %d, ii: %d\n", i, ii)
	i--
	digui(i)
}

func main() {
	digui(4)
}
