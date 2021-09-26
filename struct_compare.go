package main

import (
    "fmt"
)

type a struct {
    name string
    no int
}

func main() {
   a1 := &a{
      name: "123",
      no: 666,
   } 

   a2 := &a{
      name: "123",
      no: 666,
   }

   fmt.Println(*a1 == *a2)
}
