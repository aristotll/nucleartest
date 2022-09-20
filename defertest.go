package main

import (
	"fmt"
)

// defer 在 for 循环中的使用

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println(i)
		defer func() {
			fmt.Println("in defer: ", i)
		}()
	}
}
