package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	unicode.IsDigit('1')
	unicode.IsSpace(' ')
	if false {
		fmt.Println("123")
	}

	s := "    "
	fmt.Println(s == "")

	s1 := "       1       "
	s2 := strings.TrimSpace(s1)
	fmt.Printf("s2: %v\n", s2)

	s3 := strings.TrimSpace(s)
	fmt.Printf("s3: %v\n", s3)

	s4 := "1"
	unicode.IsDigit(rune(s4))
}
