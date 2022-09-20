package main

import "fmt"

type Point struct {
	X, Y int
}

type Rect struct {
	X, Y, W, H int
}

type Elli struct {
	X, Y, W, H int
}

type Interface interface {
	Point | Rect | Elli
}

func fn[T int | int64 | string](t T) {
	fmt.Println(t)
}

func main() {
	fn(1)
	fn("5")
	fn('1')
}
