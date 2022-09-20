package main

import (
	"fmt"
)

type Interface interface {
	xxx()
}

type S struct{}

func (s *S) xxx() {}
func (s *S) yyy() { fmt.Println("S.yyy") }

type SS struct{}

func (s *SS) xxx() {}
func (s *SS) zzz() { fmt.Println("SS.zzz") }

func fn(i Interface) {
	switch i.(type) {
	case *S:
		i.(*S).yyy()
	case *SS:
		i.(*SS).zzz()
	}
}

func main() {
	fn(&S{})
	fn(&SS{})
}
