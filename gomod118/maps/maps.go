package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3

	keys := maps.Keys(m)
	fmt.Println(keys)

	vals := maps.Values(m)
	fmt.Println(vals)

	maps.Clear(m)
	fmt.Println(m)
}
