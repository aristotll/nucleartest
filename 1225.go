package main

import (
	"fmt"
)

type People struct{}

func (p *People) A() {
	fmt.Println("people A")
	p.B()
}

func (p *People) B() {
	fmt.Println("people B")
}

type Teacher struct {
	People
}

func (t *Teacher) B() {
	fmt.Println("teacher B")
}

func main() {
	t := Teacher{}
	t.A()
}
