package main

import (
	"fmt"
)

type H map[string]interface{}

type Info struct {
	info H
}

var a = func(a int) {
	fmt.Println("param is :", a)
	h := H{}
	h["a"] = 123
	h["b"] = "bbac"

	i:= &Info{info: H{"123": "3124"}}

	fmt.Printf("%T %+v \n", i, *i)
	fmt.Println(h)

	for k, v := range i.info {
		fmt.Printf("K :%s V :%s \n", k, v)
	}
}

func test(f func(a int)) {
	f(5)
}

func main() {
	test(a)
}
