package main

import (
	"fmt"
)

type st struct {
	x, y int64
}

var m = map[string]st{}

func do(key string) {
	v, ok := m[key]
	if !ok {
		return
	}
	v.x = 9999
	fmt.Printf("[func] v: %v\n", v)
}

func main() {
	s := st{x: 10, y: 20}
	k := "a"
	m[k] = s
	do(k)
	fmt.Printf("[main] after call, s: %v\n", s)
}
