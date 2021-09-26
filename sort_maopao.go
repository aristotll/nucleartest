package main

import (
	"fmt"
)

func main() {
	n := []int{124, 1, 34, 999, 0, -1}

	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
		fmt.Println(n)
	}

	fmt.Println(n)
}
