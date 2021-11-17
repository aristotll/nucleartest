package main

import (
    "sort"
    "fmt"
)

func main() {
    s := []string{"user", "admin"}
    sort.Strings(s)

    r := sort.SearchStrings(s, "user")
    fmt.Println(r)

    r = sort.SearchStrings(s, "admin")
    fmt.Println(r)

    r = sort.SearchStrings(s, "123")
    fmt.Println(r)
}
