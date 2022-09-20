package main

import (
	"fmt"
)

/**
 *
 * @param N int整型
 * @return int整型
 */
func GetCoinCount(N int) int {
	// write code here
	sub := 1024 - N
	res := 0

	v64 := sub / 64
	sub -= v64 * 64

	v16 := sub / 16
	sub -= v16 * 16

	v4 := sub / 4

	fmt.Println(v64, v16, v4)

	res = v64 + v16 + v4
	return res
}

func main() {
	r := GetCoinCount(200)
	fmt.Println(r)
}
