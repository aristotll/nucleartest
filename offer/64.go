package main

import (
	"fmt"
)

var sum = 0

func sumNums(n int) int {
	_ = n > 1 && sumNums(n-1) > 0
	sum += n
	return sum
}

func main() {
	v := sumNums(2)
	fmt.Println(v)
}
