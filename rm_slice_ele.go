package main

import (
	"fmt"
)

// 删除切片中的奇数
func main() {
	n := []int{1, 1, 3, 5, 1, 1, 2}

	for i := 0; i < len(n); i++ {
		if n[i]%2 != 0 {
			n = append(n[:i], n[i+1:]...)
			i--
		}
	}

	fmt.Println(n)
}
