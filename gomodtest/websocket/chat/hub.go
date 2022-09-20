package chat

// hub 用于维护在线用户列表以及广播消息

type Hub struct {
	// 在线用户列表
	clients map[*Client]bool
	// 接收所用用户发送的消息
	broadcast chan []byte
	// 上线
	online chan *Client
	// 下线
	offline chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		online:    make(chan *Client),
		offline:   make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		// 有用户上线了
		case client := <-h.online:
			// 将该上线用户添加到在线列表
			h.clients[client] = true
		// 有用户下线了	
		case client := <-h.offline:
			// 从在线列表中移除
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)

			}
		// 有用户发送了一条消息
		case msg := <-h.broadcast:
			// 向全体在线用户广播
			for client := range h.clients {
				client.send <- msg
			}
		}
	}
}