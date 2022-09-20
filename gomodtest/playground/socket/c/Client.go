package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func clientHandle(conn net.Conn) {
	// 疑问：为什么先执行写操作正常，而先执行读操作会阻塞?
	// 初步分析：因为服务端先执行的是读操作，如果客户端也读，
	// 则会造成双方都在等待读取，从而阻塞

	go func() {
		// ----------- 写 ---------------------
		for {
			ioRead := bufio.NewReader(os.Stdin)
			line, _, err := ioRead.ReadLine()
			if err != nil{
				fmt.Println("从控制台读取失败, err: ", err)
				return
			}
			_, err = conn.Write(line)
			// _, err = conn.Write([]byte("fku"))
			if err != nil {
				fmt.Println("向服务端发送消息失败, err: ", err)
			}
		}
	}()

	// ----------- 读 ---------------------
	const LENS = 1024
	bytes := make([]byte, LENS)

	for {
		read, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("读取服务端响应消息失败，err: ", err)
		}
		msg := string(bytes[:read])
		fmt.Println("服务端响应消息：", msg)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":10086")
	if err != nil {
		fmt.Println("连接服务器失败")
		return
	}
	fmt.Println("客户端启动...")

	clientHandle(conn)


}
