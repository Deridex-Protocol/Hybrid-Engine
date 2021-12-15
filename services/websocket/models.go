package websocket

import (
	"github.com/gorilla/websocket"
)

const (
	WsMessageSubscribeType   = "subscribe"
	WsMessageUnsubscribeType = "unsubscribe"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientRequest struct {
	Type     string   `json:"type"`
	Channels []string `json:"channels"`
}
