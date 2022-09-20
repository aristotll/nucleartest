package main

import (
    "fmt"
)

func main() {
    var x, y int
    _, _ = x, y
    var s = []int{1, 2, 3}
    var str = "123"
    var str_ = str[:1]
    fmt.Println(str_)

    //for var k, v := range s {
    //    fmt.Println(k, v)
    //}
    if s[0] == 1 {
        goto END
    }

END:
    fmt.Println("end")
}
