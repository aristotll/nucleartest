package main

import (
    "net"
    "log"
)

func main() {
    l, err := net.Listen("tcp", ":7070")
    if err != nil {
        log.Fatalln(err)
    }

    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        conn.Write([]byte("ok"))
        conn.Close()
    }
}
