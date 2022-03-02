package main

import "fmt"

func main() {
	m := make(map[int]int)
	n := []int{1, 2, 3, 4, 5}
	for i, v := range n {
		m[i] = v
	}
	for i, v := range m {
		fmt.Println(i, v)
	}
}