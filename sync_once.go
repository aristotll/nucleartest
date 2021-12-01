package main

import (
	"sync"
	"fmt"
)

type Struct struct {
	sync.Once
}

func (s *Struct) fn() {
	for i := 0; i < 100; i++ {
		s.Do(func() {
			fmt.Println(1)
		})
	}
	
}

func (s *Struct) fn1() {
	for i := 0; i < 100; i++ {
		s.Do(func() {
			fmt.Println(2)
		})
	}
}

func main() {
	s := &Struct{}
	s.fn()
	s.fn1()
}