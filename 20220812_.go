package main

import (
    "container/heap"
    "fmt"
)

type myheap []int

func (m myheap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m myheap) Len() int {
    return len(m)
}

func (m myheap) Less(i, j int) bool {
    return m[i] < m[j]
}

func (m *myheap) Push(x interface{}) {
    *m = append(*m, x.(int))
}

func (m *myheap) Pop() interface{} {
    p := (*m)[len(*m)-1]
    *m = (*m)[:len(*m)-1]
    return p
}

func main() {
    m := []int{1, 8, 6, 2, 9}
    m_ := myheap(m)
    heap.Init(&m_)
    for m_.Len() > 0 {
        fmt.Println(heap.Pop(&m_))
    }
}
