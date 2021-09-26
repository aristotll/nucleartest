package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	stack1 *list.List
	stack2 *list.List
}

func New() *MyQueue {
	return &MyQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (m *MyQueue) Push(v int) {
	m.stack1.PushBack(v)
}

func (m *MyQueue) Pop() int {
	for m.stack1.Len() != 0 {
		pop := m.stack1.Remove(m.stack1.Back()).(int)
		m.stack2.PushBack(pop)
	}
	if m.stack2.Len() != 0 {
		pop := m.stack2.Remove(m.stack2.Back()).(int)
		return pop
	} else {
		return -1
	}
}

func main() {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
