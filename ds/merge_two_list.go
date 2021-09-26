package main

import (
    "fmt"
)

type ListNode struct {
    Next *ListNode
    Val  int
}

func NewListNode(next *ListNode, val int) *ListNode {
    return &ListNode{
        Next: next,
        Val: val,
    }
}

func merge(l1, l2 *ListNode) {
    ll1, ll2 := l1, l2
    
    for ll1 != nil && ll2 != nil {
        n := ll1.Next
        ll1.Next = ll2
        
        n1 := ll2.Next
        ll2.Next = n
        
        ll1 = n
        ll2 = n1
    }
}

func printList(h *ListNode) {
    hh := h
    for hh != nil {
        fmt.Printf("%v -> ", hh.Val)
        hh = hh.Next
    }
    fmt.Println()
}

func main() {
    // 1 -> 2 -> 3
    h := NewListNode(nil, 1)
    h1 := NewListNode(nil, 2)
    h2 := NewListNode(nil, 3)
    
    h.Next = h1
    h1.Next = h2

    // 4 -> 5
    k := NewListNode(nil, 4)
    k1 := NewListNode(nil, 5)

    k.Next = k1

    merge(h, k)
    printList(h)
}
