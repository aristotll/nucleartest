package main

import "fmt"

type MyCircularQueue struct {
	full bool
	queue []int
	addIndex, delIndex int
	len, cap int
}


func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		cap: k,
	}
}


func (m *MyCircularQueue) EnQueue(value int) bool {
	if m.len == m.cap {
		return false
	}
	fmt.Print(m.addIndex)
	if m.addIndex == m.len {
		m.addIndex = 0
	}
	m.queue[m.addIndex] = value
	m.len++
	m.addIndex++
	fmt.Println("[add]", m.len, m.queue)
	return true
}


func (m *MyCircularQueue) DeQueue() bool {
	if m.len == 0 {
		return false
	}
	if m.delIndex == m.len {
		m.delIndex = 0
	}
	m.queue[m.delIndex] = 0
	m.delIndex++
	m.len--
	return true
}


func (m *MyCircularQueue) Front() int {
	if m.len == 0 {
		return -1
	}
	return m.queue[0]
}


func (m *MyCircularQueue) Rear() int {
	if m.len == 0 {
		return -1
	}
	fmt.Println(m.len, m.queue)
	return m.queue[m.len-1]
}


func (m *MyCircularQueue) IsEmpty() bool {
	return m.len == 0
}


func (m *MyCircularQueue) IsFull() bool {
	return m.len == m.cap
}

func main() {
	q := Constructor(3)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
}
