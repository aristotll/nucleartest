package main

import (
    "fmt"
)

func fn(array [5]int) {
    fmt.Println(len(array))
}

func main() {
    i := [...]int{1, 2, 3, 4, 5}
    fmt.Printf("%T\n", i)
    fn(i)
}
