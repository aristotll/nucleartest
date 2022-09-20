package main

import "fmt"

// 结构体指针的拷贝
type Link struct {
	Val int
	Next *Link
}

func t1(l *Link) {
	fmt.Printf("t1 %v \n", *l)
	l = l.Next
	fmt.Println("t1", l)
}

func main() {
	l := &Link{
		Val:  12,
		Next: nil,
	}
	l.Next = &Link{
		Val:  22,
		Next: nil,
	}
	t1(l)
	fmt.Println("main", l)
	fmt.Printf("main %p \n", l)
}
