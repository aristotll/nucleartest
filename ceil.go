package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.1))  // 2
	fmt.Println(math.Ceil(1.45)) // 2
	fmt.Println(math.Ceil(1))    // 1
	fmt.Println(math.Ceil(1.5))  // 2
	fmt.Println(math.Ceil(1))    // 1
	fmt.Println(math.Ceil(1.01)) // 2
}
