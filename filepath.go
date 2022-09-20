package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
