package main

import (
	"fmt"
	"unicode"
)

func main() {
	c := '1'
	fmt.Println(unicode.IsLetter(c))
	fmt.Println(unicode.IsNumber(c))
}
