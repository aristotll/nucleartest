package main

import (
	"fmt"
)

func main() {
	var x = 10
	var y = 20
	
	f(x, y, func(x, y int) int {
		return x + y
	})

	f(x, y, (x, y) => {
		return x + y
	})
}

func f(x, y int, fn func(x, y int) int) {
	fmt.Println(fn(x, y))
}
