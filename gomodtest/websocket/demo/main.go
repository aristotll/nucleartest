package demo

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	// conn: 代表 websocket 连接
	// 通过调用 upgrader.Upgrade() 来创建一个连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// massageType 标识该条消息是二进制（BinaryMessage）还是字符（TextMessage）
		// 在 read 方法中，messageType 表示接收到的消息类型
		// p 是一个字节数组
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		// 在 write 方法中，messageType 用来指定发送消息的类型
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Fatal(err)
		}
	}
}
