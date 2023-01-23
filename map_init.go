package main

import (
    "fmt"
)

func main() {
    m := map[string]string{}
    m["name"] = "zhang3"
    fmt.Println(m)

    mm := make(map[string]string)
    mm["name"] = "zhang3"
    fmt.Println(mm)

    m = map[string]string{}
    fmt.Println(m)
    //var m1 map[string]string
    //m1["name"] = "zhang3"
    //fmt.Println(m1)
}
