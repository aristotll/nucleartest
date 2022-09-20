package main

import (
	"fmt"
)

type stu struct {
	name string
	age  int
}

func (s *stu) print() {
	fmt.Printf("%+v \n", s)
}

func main() {
	var s *stu
	s.print()

	var s1 = &stu{}
	s1.print()

	fmt.Println(s, s1)
}
