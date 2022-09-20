package main

import (
	"fmt"
)

type data struct {
	id int
	other string
}

var f = func() []*data{
	var d = []*data{
		{
			id:    1,
			other: "11",
		},
		{
			id:    0,
			other: "22",
		},
		{
			id: 12,
			other: "33",
		},
	}
	return d
}

func main() {
	var i = f()
	fmt.Printf("%T \n", i)
	for _, v := range i {
		fmt.Println(*v)
	}
}
