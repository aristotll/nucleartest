package main

import (
    "fmt"
    "strconv"
)

func main() {
    i, err := strconv.ParseInt("ff", 16, 64)
    if err != nil {
        panic(err)
    }
    //strconv.FormatInt(i, 10)
    fmt.Println(i)
}
