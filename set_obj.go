package main

import (
	"fmt"
)

type stu struct {
	age  int8
	name string
}

func (s *stu) equal(s1 *stu) bool {
	return s.age == s1.age
}

type set []*stu

func (s *set) add(ss *stu) bool {
	if len(*s) == 0 {
		*s = append(*s, ss)
		return true
	}

	for i := 0; i < len(*s); i++ {
		if !(*s)[i].equal(ss) {
			continue
		} else {
			return false
		}
	}

	*s = append(*s, ss)
	return true
}

func main() {
	s := new(set)
	ss := []*stu{
		{15, "a"},
		{15, "b"},
		{20, "c"},
		{20, "d"},
	}

	for i := 0; i < len(ss); i++ {
		s.add(ss[i])
	}

	for i := 0; i < len(*s); i++ {
		fmt.Printf("%+v \n", (*s)[i])
	}
}
