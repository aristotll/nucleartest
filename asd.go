package main

import (
    "fmt"
    "sort"
)

func main() {
    s := "frteafwt"
    s1 := "rthdfvji123"
    s = sortStr(s)
    s1 = sortStr(s1)
    fmt.Println(s)
    fmt.Println(s1)
}

func sortStr(s string) string {
    b := []byte(s)
    sort.Slice(b, func(i, j int) bool {return b[i] < b[j]})
    return string(b)
}
