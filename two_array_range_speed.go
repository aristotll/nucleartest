package main

import (
	"fmt"
	"time"
)

func fn1() {
	n := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	t := time.Now()
	for i := 0; i < len(n[0]); i++ {
		for j := 0; j < len(n); j++ {
			fmt.Print(n[j][i])
		}
	}
	fmt.Println("纵向遍历用时：", time.Since(t))
}

func fn2() {
	n := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	t := time.Now()
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n[0]); j++ {
			fmt.Print(n[i][j])
		}
	}
	fmt.Println(" 横向遍历用时：", time.Since(t))
}

func main() {
	fn1()
	fmt.Println()
	fn2()
}
