package main

import (
	"fmt"
	"strconv"
)

var p float64 = float64(10) / float64(1000)

func fn(n int) float64 {
	fmt.Printf("%.6f\n", p)
	f, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", p), 64)
	fmt.Println("f: ", f)
    return float64(n) * p
}

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(fn(n))
}