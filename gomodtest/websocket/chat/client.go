package chat

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// 超时时间
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// 发送消息的最大大小
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

type Client struct {
	conn *websocket.Conn
	hub *Hub
	// 每个用户发送的消息
	send chan []byte
}

// 读取消息到 hub
func (c *Client) read() {
	defer func() {
		c.hub.offline <- c
		c.conn.Close()
	}()
	// 设置读取消息的最大大小
	c.conn.SetReadLimit(maxMessageSize)
	// 设置读取超时时间
	c.conn.SetReadDeadline(time.Now().Add(pongWait))

	// todo 不知道干嘛的
	c.conn.SetPongHandler(func(s string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		// 去掉消息中的换行和空格
		message = bytes.TrimSpace(
			bytes.Replace(message, newLine, space, -1))
		// 将消息传给 hub 的 broadcast 属性，使消息广播给所有人
		c.hub.broadcast <- message
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newLine)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

