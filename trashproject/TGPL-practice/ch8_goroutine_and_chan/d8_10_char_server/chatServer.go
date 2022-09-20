package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 对外发送消息的通道
type Client chan<- string

// 通道里又存储了通道
var (
	// 用户进入了
	entering = make(chan Client)
	// 用户离开了
	leaving  = make(chan Client)
	// todo ????
	message  = make(chan string)
)

// todo 广播器
//
func broadcaster() {
	// 所有连接的客户端
	// Key 一个只能发送的通道，用来发送消息
	// Value 用户的在线状态
	clients := make(map[Client]bool)
	// 不断接收事件，并广播给用户
	for {
		// 获取当前进入事件，并进行相应的处理
		select {
		// 如果是有用户发送了消息
		case msg := <-message:
			// 把接收到的消息广播给所有用户
			for client := range clients {
				client <- msg
			}
		// 有用户上线了
		case client := <-entering:
			// 将其对应的 value 设置为 true
			clients[client] = true
		// 有用户离线了
		case client := <-leaving:
			// 删除 map 中的 key
			delete(clients, client)
			// 并关闭其用来发送消息的通道
			close(client)
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handlerConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	// 用户的 ip 地址
	who := conn.RemoteAddr().String()
	ch <- "your are " + who
	// 通知有用户上线了
	message <- who + "has arrived"
	//
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ":" + input.Text()
	}

	leaving <- ch
	message <- who + "has left"
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handlerConn(conn)
	}
}
