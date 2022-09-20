package main

import (
	"fmt"
)

type it interface {
	Name() string
	Age() int
}

type aaa struct {
	name string
	age  int
}

func NewAaa(name string, age int) *aaa {
	return &aaa{name, age}
}

func (a *aaa) Name() string {
	return a.name
}

func (a *aaa) Age() int {
	return a.age
}

func fn(i it) {
	fmt.Printf("[fn]%+v \n", i)
	fmt.Printf("[fn]%p \n", i)
	//fmt.Printf("[fn]%p \n", &i)
}

func main() {
	//a := &aaa{
	//	name: "zhang",
	//	age: 11,
	//}
	a := NewAaa("za", 11)
	fmt.Printf("[main]%p \n", a)

	fn(a)
}
