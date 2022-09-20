package main

import (
	"fmt"
	"sort"
)

func CheckPermutation(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)

	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b2[j]
	})
	fmt.Printf("b1: %v\n", string(b1))

	// sort.Slice(b2, func(i, j int) bool {
	//     return b2[i] < b2[j]
	// })
	// fmt.Printf("b2: %v\n", string(b2))

	// fmt.Println(string(b1), string(b2))
	return string(b1) == string(b2)
}

func main() {
	// s := "pnsavzruz"
	// b := []byte(s)

	// sort.Slice(b, func(i, j int) bool {
	// 	return b[i] < b[j]
	// })

	// s = string(b)

	// fmt.Printf("s: %v\n", s)
	CheckPermutation("asvnpzurz", "urzsapzvn")
}
