package main

import (
    "bytes"
    "math/rand"
    "sync"
    "net"
    "log"
)

var p = &sync.Pool {
    New: func() interface{} {
        b := make([]byte, 1024)
        return bytes.NewBuffer(b)
    },
}

func handler(conn net.Conn) error {
    for {
        b := p.Get().(*bytes.Buffer)
        // 因为 b 可能是复用的，会残留之前的数据，所以需要将数组中的每个元素初始化为 0
        memset(b.Bytes())
        _, err := conn.Read(b.Bytes())
        if err != nil {
            return err
        }

        _, err = conn.Write(b.Bytes())
        if err != nil {
            return err
        }

        p.Put(b) // 用完以后重新放回到池中
    }

}

func memset(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0
    }
}

func main() {
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Println(err)
            break
        }

        go func() {
            if err := handler(conn); err != nil {
                log.Println(err)
            }
        }()
    }
}
