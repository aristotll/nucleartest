package main

import (
	"fmt"
	"unicode"
)

func fn(s string) rune {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	var max rune
	var maxCount = -10000

	for _, v := range s {
		if m[v] > maxCount {
			maxCount = m[v]
			max = v
		}
	}

	return max
}

func fn1(s string) rune {
	m := make(map[rune]int)
	var max rune
	var maxCount = -10000

	for i := len(s) - 1; i >= 0; i-- {
		v := rune(s[i])
		if !unicode.IsLetter(v) {
			continue
		}
		m[v]++
		if m[v] >= maxCount {
			maxCount = m[v]
			max = v
		}
		//fmt.Println(string(max))
	}

	return max
}

func main() {
	//s := "aaaahfdfbbbbbbbbbb"
	//fmt.Println(string(fn1(s)))

	s1 := "hello world, every body!"
	fmt.Println(string(fn1(s1)))
}
