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

func topKFrequent(words []string, k int) []string {
	m := &myheap{
		mp: make(map[string]int),
	}

	for _, v := range words {
		m.mp[v]++
	}

	for v := range m.mp {
		if m.Len() < k {
			heap.Push(m, v)
		} else {
			if m.mp[v] > m.mp[m.s[0]] {
				heap.Pop(m)
				heap.Push(m, v)
			}
		}
	}
	fmt.Println(m.s)
	// 因为是小顶堆，堆顶为出现次数最小的，而题目要求次数多的排在前面，所以需要反转结果
	for i, j := 0, len(m.s)-1; i < j; i++ {
		m.s[i], m.s[j] = m.s[j], m.s[i]
		j--
	}
	//sort.Strings(m.s)
	return m.s
}

func main() {
	s := topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 3)
	fmt.Println(s)
}
