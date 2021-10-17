package main

import (
	"fmt"
	"runtime/debug"
)

func a() error {
	return b()
}

func b() error {
	debug.PrintStack()
	return fmt.Errorf("b error")
}

func main() {
	if err := a(); err != nil {
		return
	}
}
