package main

import (
	"fmt"
)

func main() {
	var r *int64
	{
		var x int64 = 5
		r = &x
	}
	fmt.Println(*r)
}
