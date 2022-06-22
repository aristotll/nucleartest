package main

import (
	"fmt"
)

type s struct {
	x, y int64
}

func fn(x, y int64) s {
	return s{x, y}
}

func main() {
	v := fn(1, 2)
	fmt.Println(&v)
	
}
