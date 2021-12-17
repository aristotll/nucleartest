package main

import (
	"fmt"
)

func main() {
	n := []int{1, 2, 3, 4, 5}
	nn := []int{1, 3, 4, 5, 6}
	fmt.Printf("%v result: %v \n", n, check(n))
	fmt.Printf("%v result: %v \n", nn, check(nn))
}

func check(n []int) bool {
	var index = -1
	for _, v := range n {
		if index == -1 || v-1 == index {
			index = v
			continue
		} 
		//if v-1 == index {
		//	index = v
		//	continue
		//}
		return false
	}
	return true
}


