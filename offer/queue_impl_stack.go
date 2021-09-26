package main

import (
	"container/list"
	"fmt"
)

// 一个队列实现
type MyStackByOne struct {
	Queue *list.List
}

func NewByOne() *MyStackByOne {
	return &MyStackByOne{
		Queue: list.New(),
	}
}

func (m *MyStackByOne) Push(v int) {
	_len := m.Queue.Len()
	m.Queue.PushBack(v)

	for i := 0; i < _len; i++ {
		m.Queue.PushBack(m.Queue.Remove(m.Queue.Front()))
	}
}

func (m *MyStackByOne) Pop() int {
	if m.Queue.Len() > 0 {
		return m.Queue.Remove(m.Queue.Front()).(int)
	}
	return -1
}

func (m *MyStackByOne) Top() int {
	if m.Queue.Len() == 0 {
		return -1
	} else {
		return m.Queue.Front().Value.(int)
	}
}

func testOne() {
	q := NewByOne()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println("stack push: 1 -> 2 -> 3")

	fmt.Println("top: ", q.Top())
	fmt.Println("pop: ", q.Pop())
	fmt.Println("pop: ", q.Pop())
	fmt.Println("pop: ", q.Pop())
	fmt.Println("pop: ", q.Pop())
	fmt.Println("top: ", q.Top())

}

// 两个队列实现
type MyStackByTwo struct {
	Queue1 *list.List
	Queue2 *list.List
}

func NewByTwo() *MyStackByTwo {
	return &MyStackByTwo{
		Queue1: list.New(),
		Queue2: list.New(),
	}
}

func (m *MyStackByTwo) Push(v int) {
	m.Queue2.PushBack(v)
	for m.Queue1.Len() > 0 {
		m.Queue2.PushBack(m.Queue1.Remove(m.Queue1.Front()))
	}
	m.Queue1, m.Queue2 = m.Queue2, m.Queue1
}

func (m *MyStackByTwo) Pop() int {
	if m.Queue1.Len() > 0 {
		pop := m.Queue1.Remove(m.Queue1.Front()).(int)
		return pop
	} else {
		return -1
	}
}

func (m *MyStackByTwo) Top() int {
	return m.Queue1.Front().Value.(int)
}

func testTwo() {
	s := NewByTwo()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("stack push: 1 -> 2 -> 3")

	fmt.Println("top: ", s.Top())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("pop: ", s.Pop())
	fmt.Println("pop: ", s.Pop())
}

func main() {
	//testOne()
	testTwo()
}
