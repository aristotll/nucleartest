package main

import (
	"strings"
	"fmt"
)

func main() {
	var sb strings.Builder
	sb.WriteString("123")
	fmt.Println(sb.String())
	
	sb.Reset()
	fmt.Println(sb.String())

	
}