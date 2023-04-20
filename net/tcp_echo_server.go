package main

import (
    "bufio"
    //"fmt"
    "net"
    "strings"
)

func main() {
    // 监听地址和端口
    ln, err := net.Listen("tcp", ":8081")
    if err != nil {
        panic(err)
    }
    defer ln.Close()

    for {
        // 接受客户端连接请求
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
        }

        go func(conn net.Conn) {
            // 创建一个 bufio 包装连接
            r := bufio.NewReader(conn)

            // 读取客户端的请求
            for {
                msg, err := r.ReadString('\n')
                if err != nil {
                    break
                }

                // 去除字符串尾部的换行符
                msg = strings.TrimRight(msg, "\n")

                // 将收到的字符串回写给客户端
                conn.Write([]byte(msg + "\n"))
            }

            // 关闭连接
            conn.Close()
        }(conn)
    }
}

