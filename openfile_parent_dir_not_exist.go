package main

import (
    "os"
)

func main() {
    fileP := "./a/b/c/d"
    _, err := os.OpenFile(fileP, os.O_CREATE|os.O_RDWR, 0777)
    if err != nil {
        panic(err)
    }
    os.Remove(fileP)
}
