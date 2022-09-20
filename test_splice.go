package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "localhost:8080"
	s2 := strings.Split(s, ":")
	fmt.Printf("s2: %v\n", s2)
}
