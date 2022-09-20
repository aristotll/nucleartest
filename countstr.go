package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

var s = flag.String("s", "", "input a string")

func main() {
	flag.Parse()
	fmt.Println(utf8.RuneCountInString(*s))
}
