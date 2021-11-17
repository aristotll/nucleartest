package main

import (
	"fmt"
)

type A interface {
	XX()
}

type B struct {}
func (b *B) XX() {}

type C struct {}

func main() {
	b := &B{}
	c := &C{}

	switch b {
	case A :
		fmt.Println("B is impl A")
	}
}
