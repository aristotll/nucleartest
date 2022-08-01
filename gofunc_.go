package main

import (
	"fmt"
)

func do(x, y int64, fn func(int64, int64) int64) int64 {
	return fn(x, y)
}

func main() {
	res := do(1, 2, func(x, y int64) int64 {return x + y})
	fmt.Println(res)
}
