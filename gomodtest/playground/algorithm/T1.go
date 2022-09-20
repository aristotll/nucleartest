package main

import "fmt"

func isDuplicated(a []int) bool {
	m := make(map[int]int)
	b := true
	for _, v := range a {
		if _, ok := m[v]; ok {
			b = false
		}
		m[v] = v
	}
	return b
}

func main() {
	a := []int{1, 2, 3, 5, 5}
	r := isDuplicated(a)
	fmt.Println(r)
}