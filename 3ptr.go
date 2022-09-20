package main

import (
    "fmt"
)

func main() {
     a := 1
     b := &a
     c := &b
     fmt.Printf("%T\n", c)

     *(*c) = 5
     fmt.Println(*b, **c)

     d := &c
     ***d = 100
     fmt.Println(*b, a)
}
