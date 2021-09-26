package main

import (
    "fmt"
)

type Stack []int

func NewStack() *Stack{
    s := make(Stack, 0)
    //ss := Stack(s)
    return &s
}

func (s *Stack) Push(val int) {
    *s = append(*s, val)
}

func (s *Stack) Pop() int {
    top := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return top
}

func (s *Stack) Top() int {
    return (*s)[len(*s)-1]
}

func main() {
    s := NewStack()
    s.Push(1)
    s.Push(2)
    top := s.Top()
    fmt.Println("top: ", top)
    pop := s.Pop()
    fmt.Println("pop: ", pop)
    fmt.Println("after pop: ", s)
}
