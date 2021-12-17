package main

import (
	"fmt"
)

type Struct struct {
	int64
}

func main() {
	s := &Struct{int64: 123}
	fmt.Printf("%+v \n", s)
}
