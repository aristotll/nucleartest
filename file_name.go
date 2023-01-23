package main

import (
    "os"
    "fmt"
    "path/filepath"
)

func main() {
    p := filepath.Join(os.Getenv("HOME"), "pj/justtest/file_name.go")
    f, _ := os.Open(p)
    fmt.Println(f.Name())  
}
