package main

import (
	"fmt"
)

func sort(n []int) {
	if len(n) <= 1 {
		return
	}

	p := n[0]
	l := 0
	r := len(n) - 1

	for l < r {
		for l < r && n[r] > p {
			r--
		}
		n[l] = n[r]

		for l < r && n[l] <= p {
			l++
		}
		n[r] = n[l]
	}
	fmt.Printf("%v, p = %v \n", n, p)
	n[l] = p
	sort(n[:l])
	sort(n[l+1:])
}

func main() {
	n := []int{5, 20, 13, -1, 3, 22}
	sort(n)
	fmt.Println(n)
}
