package main

import (
	"fmt"
)

func main() {
	y := {
		x := 1
		return x++
	}
	fmt.Println(y)
}
