package main

import (
    "encoding/binary"
    "fmt"
    "bytes"
)

func main() {
    //b := []byte("123")
    b := "123456"
    buf := new(bytes.Buffer)
    bb := make([]byte, len(b), len(b))

    if err := binary.Write(buf, binary.BigEndian, b); err != nil {
        panic(err)
    }

    if err := binary.Read(buf, binary.BigEndian, bb); err != nil {
        panic(err)
    }

    fmt.Println(b, bb)
}
