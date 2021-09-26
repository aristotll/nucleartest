package main

import (
	"fmt"
)

func merge(l, r []int) {
	var (
		n []int
		i = 0
		j = 0
	)

	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			n = append(n, l[i])
			i++
		} else {
			n = append(n, r[j])
			j++
		}
	}
	n = append(n, l[i:]...)
	n = append(n, r[j:]...)

	copy(l, n[:len(l)])
	copy(r, n[len(l):])
}

func sort(n []int) []int {
	if len(n) < 2 {
		return n
	}

	mid := len(n) >> 1

	l := sort(n[:mid])
	r := sort(n[mid:])

	fmt.Printf("before: %v %v ", l, r)
	merge(l, r)
	fmt.Printf("after: %v %v \n", l, r)
	
	return n
}

func main() {
	n := []int{-100, 5, 20, 13, -1, 3, 22}
	sort(n)
	fmt.Println(n)
}
