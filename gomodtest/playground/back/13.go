package main

import "fmt"

func main() {
	n := 10
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i * (i+1))
	}
	fmt.Println(sum)
}
