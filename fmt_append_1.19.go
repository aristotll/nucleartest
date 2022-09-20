package main

import (
    "fmt"
)

func main() {
    b := []byte("hello")
    b = fmt.Append(b, ", world ")
    fmt.Println("1: ", string(b))

    b = fmt.Appendln(b, "666")
    fmt.Println("2: ", string(b))

    b = fmt.Appendf(b, "%v=(%v)", "name", 123)
    fmt.Println("3: ", string(b))
}
