package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
)

var input = flag.String("s", "", "input")

func genMd5(s string) [16]byte {
    flag.Parse()
    if *input == "" {
        fmt.Println("usage: -s [input]")
        return [16]byte{}
    }

    h := md5.New()
    io.WriteString(h, *input)
    
    var r *[16]byte
    var hr = h.Sum(nil)
    r = (*[16]byte)(hr)
    return *r
}

type S struct {
    V [16]byte
}

func main() {
    var (
        buf = new(bytes.Buffer)
        s = &S{V: genMd5("123")}
        s_ = &S{}
    )
    binary.Write(buf, binary.BigEndian, s)
    binary.Read(buf, binary.BigEndian, s_)
    fmt.Printf("%x\n", s_.V)
}
