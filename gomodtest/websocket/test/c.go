package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("错误信息:", err)
	}
	wg.Add(2)
	go read(conn)
	go writeM(conn)
	wg.Wait()
}

func read(conn *websocket.Conn) {
	defer wg.Done()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("错误信息:", err)
			break
		}
		if err == io.EOF {
			continue
		}

		fmt.Printf("<%s> %s\n", conn.RemoteAddr().String(), string(msg))
	}
}

func writeM(conn *websocket.Conn) {
	defer wg.Done()
	for {
		//fmt.Print("请输入:")
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		if err := conn.WriteMessage(websocket.TextMessage, []byte(data)); err != nil {
			log.Println("conn write error: ", err)
			return
		}
	}
}
