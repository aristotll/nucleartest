package main

import (
	"flag"
	"fmt"
	"strings"
)

type StringSlice []string

func (s *StringSlice) String() string {
	return fmt.Sprintf("%v", StringSlice(*s))
}

func (s *StringSlice) Set(value string) error {
	ss := strings.Split(value, ",")
	*s = append(*s, ss...)
	return nil
}

func main() {
	var s StringSlice
	flag.Var(&s, "s", "slice")
	flag.Parse()
	fmt.Println(s)
}
