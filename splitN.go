package main

import (
	"strings"
	"fmt"
	"flag"
)

var (
	str = flag.String("s", "", "input a string")
	sep = flag.String("p", "", "分隔符")
	n = flag.Int("n", 0, "input N")
)

func main() {
	flag.Parse()
	r := strings.SplitN(*str, *sep, *n)
	fmt.Printf("SplitN(%v, %v, %v) = %v\n", *str, *sep, *n, r)
}
