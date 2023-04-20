package main

import (
    "fmt"
    "net/rpc"
)

type Struct struct {}

type Req struct {}

type Resp struct {
    R []string
}

func main() {
    cli, err := rpc.Dial("tcp", ":9090")
    if err != nil {
        panic(err)
    }
    for {
        resp := new(Resp)
        if err := cli.Call("Struct.Do", &Req{}, resp); err != nil {
            fmt.Println(err)
            continue
        }
        fmt.Println(resp)
    }
}
