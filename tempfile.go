package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	f, err := ioutil.TempFile("testgeneric", "example")
	if err != nil {
		panic("create temp file error: " + err.Error())
	}
	fmt.Println(f.Name())
}
