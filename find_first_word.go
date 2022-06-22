package main

import (
	"fmt"
)

func findFirstWord(s string) (first string) {
	for _, char := range s {
		if string(char) == " " {
			break
		}
		first += string(char)
	}
	return
}

func main() {
	fmt.Println(findFirstWord("My name is"))
}
