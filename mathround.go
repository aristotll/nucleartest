package main

import (
	"fmt"
	"math"
)

func P(v float64) {
	fmt.Println(math.Round(v))	
}

func main() {
	P(1.5)
	P(1.15)
	P(1)
}
