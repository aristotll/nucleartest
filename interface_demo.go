package main

import (
	"fmt"
)

type P interface {
	Work()
}

type PA struct {
}

func (p *PA) Work() {
	fmt.Println("pa word!")
}

type A struct {
	P
}

func main() {
	a := &A{
		P: &PA{},
	}
	a.Work()
}
