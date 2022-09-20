package main

import (
	"container/list"
	"fmt"
)

func exchange(l1, l2 *list.List) {
	*l1, *l2 = *l2, *l1
}

func exNo(l1, l2 *list.List) {
	l1, l2 = l2, l1
}

func main() {
	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	l2 := list.New()
	l2.PushBack(666)

	fmt.Println("before: ")
	_print("l1", l1)
	_print("l2", l2)

	exchange(l1, l2)
	//exNo(l1, l2)
	fmt.Println("\nafter: ")
	_print("l1", l1)
	_print("l2", l2)

	fmt.Println()
	l2.PushBack(4)
	_print("l2", l2)
}

func _print(name string, l *list.List) {
	fmt.Printf("%v : ", name)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v -> ", i.Value)
	}
}
