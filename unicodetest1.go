package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "a1b2"
	b := []byte(s)
	var res []string

	for i := 0; i < len(b); i++ {
		r := rune(b[i])
		if unicode.IsLetter(r) {
			lower := unicode.ToLower(r)
			upper := unicode.ToUpper(r)

			b[i] = byte(lower)
			res = append(res, string(b))

			b[i] = byte(upper)
			res = append(res, string(b))
		}
	}

	fmt.Println(res)
}
