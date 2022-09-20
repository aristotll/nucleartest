package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var mux sync.Mutex
var __wg__ sync.WaitGroup

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var hub = NewHub()

type Client struct {
	conn    *websocket.Conn
	message chan []byte
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		case msg := <-h.broadcast:
			for client := range h.clients {
				client.message <- msg
			}
		}
	}
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 连接成功后注册用户
	client := &Client{
		conn:    conn,
		message: make(chan []byte),
	}
	defer func() {
		hub.unregister <- client
		addr := client.conn.RemoteAddr().String()
		smsg := fmt.Sprintf("<%s>离线了", addr)
		hub.broadcast <- []byte(smsg)
	}()
	// 添加到注册用户中
	hub.register <- client
	// 广播上线信息
	addr := client.conn.RemoteAddr().String()
	smsg := fmt.Sprintf("<%s>上线了", addr)
	hub.broadcast <- []byte(smsg)

	__wg__.Add(1)
	go client.read()
	go client.write()
	__wg__.Wait()
}

func (c *Client) read() {
	for {
		_, msg, err := c.conn.ReadMessage()
		// websocket: close 1006 (abnormal closure): unexpected EOF
		if err != nil {
			fmt.Println("server read error: ", err)
			fmt.Println("用户退出：", c.conn.RemoteAddr().String())
			hub.unregister <- c

			addr := c.conn.RemoteAddr().String()
			smsg := fmt.Sprintf("<%s>离线了", addr)
			hub.broadcast <- []byte(smsg)

			break
		}
		// 将读取到的信息传入websocket处理器中的broadcast中
		hub.broadcast <- msg
	}
}

func (c *Client) write() {
	defer __wg__.Done()
	for msg := range c.message {
		mux.Lock()
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		mux.Unlock()
		if err != nil {
			fmt.Println("写入错误")
			break
		}
	}
}

func main() {
	go hub.Run()
	http.HandleFunc("/", WsHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("start server error: ", err)
		return
	}
}
