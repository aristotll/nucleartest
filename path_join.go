package main

import (
	"fmt"
	"path"
)

func main() {
	p1 := "/service/"
	p2 := "//abc"
	p3 := "78da7787f787ff97a"
	p := path.Join(p1, p2, p3)
	fmt.Println(p)

	p = path.Join("http://", "/cache/", "/a", "b")
	fmt.Println(p)
}
