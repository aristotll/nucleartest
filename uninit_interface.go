package main

import (
    "fmt"
    "reflect"
)

type I interface {}

func main() {
    var i I
    fmt.Println(i)

    typ := reflect.TypeOf(&i).Elem()
    fmt.Println(typ.Kind())
}
