package main

import (
	"fmt"
	//"strings"
)

func main() {
	fmt.Println(1 << 24)
    fmt.Println(0 << 24)
	fmt.Println(128 | 4)
    fmt.Println(1 >> 24, 1 >> 16, 1 >> 8, 1 >> 0)
    fmt.Println(^byte(255), ^byte(0))
	//fmt.Println(strings.Repeat("0", 1<<24))
}
