package main

import (
	"fmt"
)

func main() {
	x := 5
	{
		x := x * 2
		fmt.Println(x)
	}
	fmt.Println(x)
}
