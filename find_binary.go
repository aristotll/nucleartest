package main

import (
    "fmt"
    "flag"
)

var input  = flag.Int("d", 0, "input dec number")

func bin(n int) int {
    var r int
    i := 1
    for n != 0 {
        r += i * (n%2)
        i *= 10
        n /= 2
    }
    return r
}

func main() {
    //flag.Int("-i", &input, "input dec number")
    flag.Parse()
    fmt.Println("input: ", input)
    r := bin(*input)
    fmt.Printf("%v\n", r)
}
