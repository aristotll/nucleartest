package main

import "fmt"

type stack []interface{}

func NewStack() *stack {
	s := make(stack, 0)
	return &s
}

func (s *stack) Push(v interface{}) *stack {
	S := s
	*S = append(*S, v)
	return S
}

func (s *stack) Range() {
	fmt.Println(s)
}

func (s *stack) Pop() *stack {
	S := s
	l := len(*s)-1
	*S = append((*S)[0:l])
	return S
}

func main() {
	s := NewStack()

	s.Push(113).Push(266).Push("dad").Push("uopty[u")
	s.Pop()
	s.Range()
}
