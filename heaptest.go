package main

import "container/heap"

type ListNode struct {
	Next *ListNode
	Val int
}

type myheap []*ListNode

func (h myheap) Len() int {
    return len(h)
}

func (h myheap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h myheap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func (h *myheap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

func (h *myheap) Pop() interface{} {
    p := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return p
}


func main() {
	var h *myheap
	heap.Push(h, &ListNode{})
}