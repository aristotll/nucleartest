package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	age  int
}

func printT(s []*student, f func(name string) bool )  {
	for _, v := range s{
		if f(v.name) {
			fmt.Println("value:", v)
			fmt.Printf("v: %+v", v)
			break
		}else if !f(v.name) {
			fmt.Println("false")
		}
	}
}

func main() {
	arr := make([]*student, 10)

	arr = []*student{
		{
			name: "zhang3",
			age:  16,
		},
		{
			name: "li4",
			age: 20,
		},
	}

	printT(arr, func(name string) bool {
		if strings.HasPrefix(name, "z") {
			return true
		}
		return false
	})
}


