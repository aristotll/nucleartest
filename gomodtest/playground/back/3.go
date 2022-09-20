package main

import "fmt"

func hash(key string, size uint) uint {
	var h uint
	for _, char := range key {
		fmt.Println("char: ", char)
		h = (h << 5) + uint(char+1)
	}
	return h % size
}

func main() {
	s := "abc"
	l := len(s)
	h := hash(s, uint(l))
	fmt.Println(h)
}
