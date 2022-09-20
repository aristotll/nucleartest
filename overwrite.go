package main

import (
    "fmt"
)

type S struct{}

func (s *S) Do() { fmt.Println("father") }

type C struct{
    S
}

func (c *C) Do() { fmt.Println("child") }

func main() {
    c := new(C)
    c.Do()
}
