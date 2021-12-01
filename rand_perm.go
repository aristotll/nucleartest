package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ints := rand.Perm(100)
	fmt.Println(ints)
}