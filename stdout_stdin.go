package main

import (
   // "fmt"
    "os"
    "io"
)

func main() {
    io.Copy(os.Stdout, os.Stdin)
}
