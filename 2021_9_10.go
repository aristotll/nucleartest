package main

import "fmt"

type cul func(x, y int64) int64

type Cli struct {
	cul
}

func NewCli(c cul) *Cli {
	return &Cli{
		cul: c,
	}
}

func (c cul) Run(x, y int64) {
	fmt.Println(c(x, y))
}

func main() {
	c := NewCli(func(x, y int64) int64 { return x + y })
	c.Run(1, 2)
}
