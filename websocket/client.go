package websocket

import (
	"log"
	"net/http"
	"time"

	ws "github.com/gorilla/websocket"
)

// WSClient is the connected client
type WSClient struct {
	hub  *Hub
	conn *ws.Conn
	send chan []byte
}

var updrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// check client is valid origin or not, currently do nothing
		return true
	},
}

var (
	newline = []byte{'\n'} // []byte("\n")
	space   = []byte{' '}  // []byte(" 00")
)

// read message from websocket connection peer and push to hub broadcast
// encapsulation message to struct Message includes wsclient info, thus
// hub could identify which client sent message, and not send duplicate message to the client.
func (c *WSClient) read() {
	defer func() {
		c.hub.unRegister <- c
		_ = c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			// close connection
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway, ws.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		if msg := string(message[:]); msg == "ping" {
			_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
			continue
		}
		// no need to share message from peer to peer in current scope
		// clientMsg := Message{Sender: c, Msg: message}
		// c.hub.broadcast <- clientMsg
	}
}

// write message to client
func (c *WSClient) write() {
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				log.Printf("client's send is closed")
			}
			w, err := c.conn.NextWriter(ws.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				log.Println(err)
			}
			//  Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-c.send)
			}
			if err = w.Close(); err != nil {
				return
			}
		case <-ticker.C:

		}
	}
}

// ServeWS handle upgrade http connection to websocket
func ServeWS(w http.ResponseWriter, r *http.Request, h *Hub) {
	conn, err := updrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &WSClient{
		conn: conn,
		hub:  h,
		send: make(chan []byte, 256),
	}

	h.register <- client

	go client.read()
	go client.write()
}
