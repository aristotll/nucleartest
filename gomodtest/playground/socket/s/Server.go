package main

import (
	"fmt"
	"net"
	"strings"
)

func serverHandle(l net.Conn) {
	// addr := l.RemoteAddr().String()
	// 尝试先向客户端响应一条消息
	//_, _ = l.Write([]byte(addr + "你已成功连接服务器"))

	// 每次读取的字节长度
	const LENS = 1024
	var bytes = make([]byte, LENS)

	for {
		read, err := l.Read(bytes)
		if err != nil {
			fmt.Println("服务端读取数据错误, 错误：", err)
			return
		}
		fmt.Println("var read: ", read)
		msg := string(bytes[:read])
		fmt.Println("接受到了客户端发来的消息：", msg)
		if strings.ToUpper(msg) == "Q" {
			_, err := l.Write([]byte("退出"))
			if err != nil {
				fmt.Println("响应消息失败")
				return
			}
		}else if msg == "hello" {
			_, err := l.Write([]byte("你好"))
			if err != nil {
				fmt.Println("响应消息失败, 错误：", err)
				return
			}
		}else {
			_, err := l.Write([]byte("我读到了你发来的消息: " + msg))
			if err != nil {
				fmt.Println("响应消息失败, err: ", err)
				return
			}
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("监听端口失败, 错误：", err)
	}
	fmt.Println("服务端启动...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败, 错误：", err)
		}
		serverHandle(conn)
	}
}
