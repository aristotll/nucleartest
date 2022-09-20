package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(strconv.FormatInt(int64(255), 2))
    fmt.Println(strconv.FormatInt(int64(255), 8))
    fmt.Println(strconv.FormatInt(int64(255), 16))
}
