package main

import "fmt"

// go run -gcflags=-G=3 stack/stack.go

// internal compiler error: Cannot (yet) export a generic type: Stack
// 不支持可导出类型的泛型
type stack[T any] struct {
	val []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(x T) {
	s.val = append(s.val, x)
}

func (s *stack[T]) Pop() T {
	p := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return p
}

func (s *stack[T]) Peek() T {
	return s.val[len(s.val)-1]
}

func main() {
	s := newStack[int]()
	s.Push(1)
	s.Push(2)

	fmt.Println(s.val)
}