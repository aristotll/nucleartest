package main

import (
    "flag"
    "fmt"
)

var b = flag.Bool("b", false, "if use mac k8s linux k3s")

func main() {
    flag.Parse()
    fmt.Println(*b)
}
