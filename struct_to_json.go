package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type st struct {
		A int64
		B string
		C struct {
			Name string
		}
	}

	s := &st{
		A: 100,
		B: "abc",
		C: struct{Name string}{"zhang3"},
	}

	b, err := json.Marshal(&s)
	if err != nil {
		panic(err)
	}

	json.MarshalIndent()

	fmt.Println(string(b))
}
