package main

import (
	"fmt"
)

func main() {
	var a uint16 = 94
	s := string(a)
	fmt.Println(s)

	s1 := fmt.Sprint(a)
	fmt.Println(s1)
}
