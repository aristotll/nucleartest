package main

import (
	"encoding/base64"
	"fmt"
	"strings"
    "flag"
)

var input = flag.String("s", "", "input")

func main() {
    flag.Parse()
    if *input == "" {
        fmt.Println("usage: -s [input]")
        return
    }

	var ss strings.Builder

	enc := base64.NewEncoder(base64.StdEncoding, &ss)
	enc.Write([]byte(*input))

	fmt.Println(ss.String())
}
