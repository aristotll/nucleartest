package main

import (
	"container/heap"
	"fmt"
)

type myheap struct {
	s  []string
	mp map[string]int
}

func (m myheap) Swap(i, j int) {
	m.s[i], m.s[j] = m.s[j], m.s[i]
}

func (m myheap) Len() int {
	return len(m.s)
}

func (m myheap) Less(i, j int) bool {
	o1 := m.s[i]
	o2 := m.s[j]
	m1 := m.mp[o1]
	m2 := m.mp[o2]
	if m1 == m2 {
		return o1 > o2 // 如果出现频率相同，按字典序排列
	}
	return m1 < m2
}

func (m *myheap) Push(x interface{}) {
	m.s = append(m.s, x.(string))
}

func (m *myheap) Pop() interface{} {
	p := m.s[len(m.s)-1]
	m.s = m.s[:len(m.s)-1]
	return p
}

func topKFrequent(words []string, k int) (res []string) {
	m := &myheap{
		mp: make(map[string]int),
	}

	for _, v := range words {
		m.mp[v]++
	}

	for v := range m.mp {
		heap.Push(m, v)
		if m.Len() > k {
			heap.Pop(m)
		}
		fmt.Println(m.s)
	}

	for m.Len() > 0 {
		v := heap.Pop(m)
		fmt.Println("pop: ", v)
	}
    return m.s
}

func main() {
	topKFrequent([]string{"i", "love", "leetcode"}, 3)
}
