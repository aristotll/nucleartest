package main

import (
	"bufio"
	"fmt"
	"os"
)

func pt() {
	for i := 0; i < 10; i++ {
		v := 100
		v++
		fmt.Println(v)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 保存进位
	temp := 0
	result := new(ListNode)

	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		sum := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum = v1 + v2 + temp
		if sum >= 10 {
			sum %= 10
			temp = 1
		}
		result = &ListNode{
			Val:  sum,
			Next: &ListNode{},
		}
		result = result.Next
	}

	file, _ := os.OpenFile("t", os.O_RDONLY, 0744)

	bufio.NewReader(file)
	return result
}

func main() {
	pt()
}
