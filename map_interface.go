package main

import "fmt"

type I interface {
	XXX()
}

type A struct {
	I
	X int64
}

func (a A) XXX() {}

type B struct {
	I
	X int64
}

func (b B) XXX() {}

var MAP = make(map[I]bool)
var (
	_ I = A{}
	_ I = B{}
)

func main() {
	var a, b = new(A), new(B)
	a.X, b.X = 123, 456

    var x = A{X: 123}
    var y = B{X: 456}

	MAP[a] = true
	MAP[b] = true
    MAP[x] = true
    MAP[y] = true

	fmt.Println(MAP)
}
