package main

import (
	"fmt"
)

func a() {defer fmt.Println("a")}
func b() {
	fmt.Println("b")
	a()
}
func c() {b()}

func main() {
	c()
}
