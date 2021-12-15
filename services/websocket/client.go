package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Client is a struct that can be used by many channels to write
type Client struct {
	ID   string
	conn *websocket.Conn
	mu   sync.Mutex
}

func NewClient(id string, conn *websocket.Conn) Client {
	return Client{
		ID:   id,
		conn: conn,
	}
}

func (c *Client) WriteJSON(message interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.conn.WriteJSON(message)
}
