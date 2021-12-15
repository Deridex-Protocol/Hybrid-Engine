package websocket

import (
	"context"
	"sync"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"github.com/sirupsen/logrus"
)

type Channel struct {
	MarketID     string
	MessagesChan chan *common.WebSocketMessage
	Clients      map[string]*Client
	clientsMu    sync.Mutex
	log          *logrus.Entry
}

func NewChannel(marketID string, log *logrus.Entry) *Channel {
	return &Channel{
		MarketID:     marketID,
		MessagesChan: make(chan *common.WebSocketMessage),
		Clients:      make(map[string]*Client),
		log:          log,
	}
}

func (c *Channel) Run(ctx context.Context) {
	for {
		select {
		case msg := <-c.MessagesChan:
			c.clientsMu.Lock()
			for clientID := range c.Clients {
				if err := c.Clients[clientID].WriteJSON(msg.Data); err != nil {
					c.log.WithError(err).
						WithField("client_id", clientID).
						Error("Failed to send message to client")
					delete(c.Clients, clientID)
					continue
				}

				c.log.WithField("client_id", clientID).Info("Send web socket message to client")
			}
			c.clientsMu.Unlock()
		case <-ctx.Done():
			c.log.Info("Stop channel")
		}
	}
}

func (c *Channel) AddClient(client *Client) {
	c.clientsMu.Lock()
	c.Clients[client.ID] = client
	c.clientsMu.Unlock()
}

func (c *Channel) RemoveClient(clientID string) {
	c.clientsMu.Lock()
	delete(c.Clients, clientID)
	c.clientsMu.Unlock()
}
