package main

import (
    "net/http"
    "io"
    "fmt"
)

type Resp struct {
    r io.Reader
}

func getr() *Resp {
    resp, err := http.Get("http://www.baidu.com")
    if err != nil {
        panic(err)
    }
    res := &Resp{}
    res.r = resp.Body
    defer resp.Body.Close()
    return res
}

func main() {
    rs := getr()
    b, err := io.ReadAll(rs.r)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}
