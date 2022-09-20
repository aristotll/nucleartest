package main

import (
	"fmt"
	set "github.com/deckarep/golang-set"
)

type stu struct {
	name string
	age  int16
}

func main() {
	s := set.NewSet()
	s.Add("1")
	s.Add("2")
	fmt.Println(s)

	s = set.NewSet()
	s.Add(stu{"123", 10})
	s.Add(stu{"123", 10})
	fmt.Println(s)
}
