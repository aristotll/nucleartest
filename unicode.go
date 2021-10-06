package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "xiaoming is a good student"
	s1 := "123 456 789"
	
	// for _, c := range s {
	// 	b := unicode.IsLetter(c)
	// 	fmt.Printf("%v: %v\n", string(c), b)
	// }

	for _, v := range strings.Fields(s1) {
		fmt.Printf("v: %v\n", v)
	}

	ss := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	for _, v := range ss {
		fmt.Printf("v: %v\n", v)
	}
}
