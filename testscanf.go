package main

import "fmt"

func main() {
	n, err := fmt.Scan("a")
	if err != nil {
		panic(err)
	}
	fmt.Printf("n: %v\n", n)
}
