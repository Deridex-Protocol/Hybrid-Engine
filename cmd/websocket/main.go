package main

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/redis"
	"bitbucket.ideasoft.io/dex/dex-backend/common/shutdown"
	"bitbucket.ideasoft.io/dex/dex-backend/services/websocket"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := shutdown.GracefulShutdown()

	var cfg websocket.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal("Failed to parse envs to struct")
	}

	log := logrus.New()
	logger := log.WithField("service", "websocket")

	redisClient, err := redis.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		logrus.Fatal("Failed to connect to redis")
	}

	service := websocket.NewWebsocketService(logger, &cfg, redisClient)
	service.Run(ctx)
}
