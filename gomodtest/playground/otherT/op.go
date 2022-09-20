package main

import "fmt"

func breakFor(a, b []int) {
	l: fmt.Println("aa")
	for _, v := range a {
		fmt.Printf("[%v]\n", v)
		for k, v := range b {
			if k == 1 {
				goto l
			}
			fmt.Printf("%v \t", v)
		}
		fmt.Println()
	}
}

func main() {
	a := make([]int, 0)
	a = append(a, 123, 43, 67)

	b := make([]int, 0)
	b = append(b, 12, 23)

	breakFor(a, b)
}
