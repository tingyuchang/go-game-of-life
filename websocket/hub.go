package websocket

type Hub struct {
	// store clients
	clients map[*WSClient]bool
	// broadcast msg to registered client
	broadcast chan Message
	// received register client
	register chan *WSClient
	// received unregister client
	unRegister chan *WSClient
}

// Run is running a for-loop to receive hub's channels
// register will store connecting client
// unregister will remove disconnecting client
// broadcast is response for broadcasting message to all registered clients
func (h *Hub) Run() {
	for {
		select {
		// register client
		case client := <-h.register:
			h.clients[client] = true
		// unregister client
		case client := <-h.unRegister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				// skip duplicate message
				if message.Sender != client {
					select {
					case client.send <- message.Msg:
					default:
						delete(h.clients, client)
						close(client.send)
					}
				}
			}
		}
	}
}

// BroadcastMsg prepare a defined Message format
// Sender is nil means this message is sent by System
func (h *Hub) BroadcastMsg(msg []byte) {
	serverMessage := Message{Sender: nil, Msg: msg}
	h.broadcast <- serverMessage
}

// NewHub create a new hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*WSClient]bool),
		broadcast:  make(chan Message),
		register:   make(chan *WSClient),
		unRegister: make(chan *WSClient),
	}
}
// Init initial currentHub
func Init() {
	CurrentHub = NewHub()
}

