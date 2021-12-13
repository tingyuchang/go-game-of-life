package websocket

type Message struct {
	Sender *WSClient
	Msg    []byte
}
