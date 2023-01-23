package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		b := make([]byte, 10)
		_, err := rand.Read(b)
		assert(err, "rand.Read error: ")
		fmt.Println(string(b))
	}
}

func assert(err error, msg string) {
	if err == nil {
		return
	}
	panic(fmt.Sprintf("%s:%s", msg, err.Error()))
}
