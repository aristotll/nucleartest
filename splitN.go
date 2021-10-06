package main

import (
	"strings"
	"fmt"
)

func main() {
	n := strings.SplitN("name/123", "/", 2)
	fmt.Println(n, len(n))
}
