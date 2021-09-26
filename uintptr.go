package main

import "fmt"

func main() {
	i := 100
	fmt.Printf("%p \n", &i)
	fmt.Println(uintptr(i))
}
