package main

import (
	"fmt"
)

var m = map[string]string{
	"name": "zhang3",
	"age":  "20",
}

func main() {
	var key = m["a"]
	fmt.Printf("%q %v\n", key, key == "")
}
