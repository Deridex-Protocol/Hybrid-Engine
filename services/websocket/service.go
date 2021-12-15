package websocket

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log      *logrus.Entry
	cfg      *Config
	redis    *redis.Client
	channels map[string]*Channel
}

func NewWebsocketService(log *logrus.Entry, cfg *Config, redis *redis.Client) Service {
	return Service{
		log:      log,
		cfg:      cfg,
		redis:    redis,
		channels: make(map[string]*Channel),
	}
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Websocket service start")

	for _, channelID := range s.cfg.Channels {
		s.channels[channelID] = NewChannel(channelID, s.log.WithField("channel_id", channelID))
		s.log.WithField("channel_id", channelID).Info("Market start")
	}

	go s.startEventConsumer(ctx)

	srv := &http.Server{Addr: ":3002"}

	http.HandleFunc("/ws", s.websocketHandler)

	go func() {
		s.log.Info("Websocket Server is listening on :3002")
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			s.log.WithError(err).Fatal("Websocket serve exit error")
		}
	}()

	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		s.log.WithError(err).Fatal("Failed shutdown websocket server")
	}

	s.log.Info("Websocket service stop")
}

func (s *Service) startEventConsumer(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			s.log.Info("Websocket consumer exit")
			return
		default:
			res, err := s.redis.BRPop(time.Second, common.WebsocketMessageQueueKey).Result()
			if err != nil && err != redis.Nil {
				s.log.WithError(err).Error("Failed to BRPop message from event queue")
				return
			}

			if err == redis.Nil {
				continue
			}

			var webSocketMessage common.WebSocketMessage
			if err = json.Unmarshal([]byte(res[1]), &webSocketMessage); err != nil {
				s.log.WithError(err).Error("Failed unmarshal message to struct")
				continue
			}

			if channel, ok := s.channels[webSocketMessage.ChannelID]; ok {
				channel.MessagesChan <- &webSocketMessage
			}
		}
	}
}

func (s Service) websocketHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Origin, Jwt-Authentication, Hydro-Authentication")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.log.WithError(err).Error("Failed to upgrade connection")
		return
	}

	client := NewClient(uuid.NewV4().String(), conn)

	s.log.WithField("client_id", client.ID).
		WithField("ip_addr", conn.RemoteAddr()).
		Info("New client connected")

	s.handleClientRequest(&client)

	s.log.WithField("client_id", client.ID).
		WithField("ip_addr", conn.RemoteAddr()).
		Info("New client disconnected")

	if err := conn.Close(); err != nil {
		s.log.WithError(err).Error("Failed to close websocket connection")
		return
	}
}

func (s *Service) handleClientRequest(client *Client) {
	for {
		var req ClientRequest
		client.mu.Lock()
		if err := client.conn.ReadJSON(&req); err != nil {
			client.mu.Unlock()
			switch err.(type) {
			case *json.SyntaxError:
				continue
			default:
				return
			}
		}
		client.mu.Unlock()

		switch req.Type {
		case WsMessageSubscribeType:
			for _, id := range req.Channels {
				if channel, ok := s.channels[id]; ok {
					channel.AddClient(client)
				}
			}
		case WsMessageUnsubscribeType:
			for _, id := range req.Channels {
				if channel, ok := s.channels[id]; ok {
					channel.RemoveClient(client.ID)
				}
			}
		}
	}
}
