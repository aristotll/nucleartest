package main

import (
	"strings"
	"fmt"
	"unicode"
)

func main() {
	s := "my name is a,   and her name is b."
	
	letters := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	fmt.Println(letters)
}