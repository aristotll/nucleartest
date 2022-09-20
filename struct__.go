package main

import (
    "fmt"
)

type People struct {
    Name string
    Age int8    
}

func NewPeople(name string, age int8) *People {
    return &People{
        Name: name,
        Age: age,
    }
}

func main() {
    p := &People{
        Name: "zhang3",
        Age: 12,
    }

    fmt.Printf("%+v\n", p)

    p = NewPeople("zhang3", 12)
    fmt.Printf("%+v\n", p)
}


