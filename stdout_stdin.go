package main

import (
	// "fmt"
	"io"
	"os"
)

func main() {
	io.Copy(os.Stdout, os.Stdin)
}
