package main

import (
	"fmt"
)

type Class struct {
	fn func(int, int) int
}

func (c *Class) DoSafe(x, y int) {
	if c.fn != nil {
		fmt.Println(c.fn(x, y))
	} else {
		fmt.Println("fn is nil")
	}
}

func (c *Class) Do(x, y int) {
	fmt.Println(c.fn(x, y))
}

func main() {
	c := &Class{fn: func(x, y int) int { return x + y }}
	c.Do(1, 2)

	cc := &Class{}
	cc.DoSafe(1, 2)
	cc.Do(1, 2)
}
