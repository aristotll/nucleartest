package main

import (
	"fmt"
)

func main() {
	type a struct {
		a, b int64
	}

	a1 := &a{1, 2}
	a2 := &a{1, 2}

	fmt.Println(a1 == a2)
	fmt.Println(*a1 == *a2)
}
