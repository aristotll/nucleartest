package main

import (
	"fmt"
)

func main() {
	//x := (y := 1)
	//fmt.Println(x) //error
	var x, y int64
	x = y = 123
	fmt.println(x, y)
}
