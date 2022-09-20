package main

import (
	"fmt"
)

type (
	Color1 int32
	Color2 int32
	Color3 int32
)

type Color struct {
	Color1
	Color2
	Color3
}

func main() {
	c := &Color{0, 0, 0}
	fmt.Println(c)
}
