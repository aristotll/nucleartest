package main

import "fmt"

type Pair[X, Y any] struct {
	Key X
	Val Y
}

func (p *Pair[X, Y]) Print() {
	fmt.Printf("key: %v, value: %v\n", p.Key, p.Val)
}

func (p *Pair[X, Y]) MethodGenerice[T any]() {
	
}

func main() {
	var p Pair[int, int]
	p.Key = 1
	p.Val = 123
	p.Print()

	var pp Pair[string, any]
	pp.Key = "666"
	pp.Val = 123
	pp.Val = "123"
	pp.Val = '1'
	pp.Print()
}
