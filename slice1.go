package main

import (
	"fmt"
)

type Vector []int

func NewVector() *Vector {
	v := make(Vector, 0)
	return &v
}

func (v *Vector) PushBack(val int) {
	*v = append(*v, val)
}

func (v *Vector) Pop() int {
	pop := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return pop
}

func (v *Vector) Get(index int) int {
	return (*v)[index]
}

func main() {
	v := NewVector()
	v.PushBack(5)
	v.PushBack(555)
	fmt.Println(v)
	pop := v.Pop()
	fmt.Println("pop: ", pop)
	fmt.Println(v)
}
