package main

import (
	"fmt"
	"sync"
)

type aa struct {
	m  map[int]int
	mu sync.RWMutex
}

func Newaa() *aa {
	return &aa{
		m: make(map[int]int),
	}
}

func (a *aa) t(i int) int {
	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.m[i]; !ok {
		fmt.Println("!ok")
		a.m[i]++
		return a.m[i]
	}

	return a.m[i]

}

func main() {
	a := Newaa()
	a.t(1)
}
