package main

import (
	"fmt"
)

type inter interface {
	Do(inter) string
}

type stc struct {
	str string
}

func (*stc) Do(s *stc) string {
	return  s.str
}

func main() {
	ss := &stc{"abc"}
	s1 := &stc{}
	ret := s1.Do(ss)
	fmt.Println(ret)
}


