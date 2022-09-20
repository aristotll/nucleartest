package main

import (
	"fmt"
	"log"
)

var i = 0

func x(arr []int) []int {
	lens := len(arr)
	m := lens / 2
	l := arr[:m]
	r := arr[m:]
	i++
	if i < 10 {
		log.Fatal()
	}
	return y(x(l), x(r))
}

func y(l, r []int) []int {
	fmt.Printf("l: %v, r: %v", l, r)

	return []int{}
}

func main() {
	x([]int{2, 1, 65, 3, 6, 7, 14})
}
