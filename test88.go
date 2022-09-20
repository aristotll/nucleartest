package main

import (
	"container/list"
)

type MyStack struct {
	queue *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		queue: list.New(),
	}
}

/** Push element x onto stack. */
func (m *MyStack) Push(x int) {
	queueLen := m.queue.Len()
	m.queue.PushBack(x)

	for i := 0; i < queueLen; i++ {
		pop := m.queue.Remove(m.queue.Front()).(int)
		m.queue.PushBack(pop)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (m *MyStack) Pop() int {
	pop := m.queue.Remove(m.queue.Front()).(int)
	return pop
}

/** Get the top element. */
func (m *MyStack) Top() int {
	return m.queue.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (m *MyStack) Empty() bool {
	return m.queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
