package main

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"bitbucket.ideasoft.io/dex/dex-backend/common/redis"
	"bitbucket.ideasoft.io/dex/dex-backend/common/shutdown"
	"bitbucket.ideasoft.io/dex/dex-backend/services/watcher"
	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := shutdown.GracefulShutdown()

	var cfg watcher.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal("Failed to parse envs to struct")
	}

	log := logrus.New()
	logger := log.WithField("service", "watcher")

	qi, err := models.Connect(cfg.DatabaseURL)
	if err != nil {
		logrus.Fatal("Failed to connect to database")
	}

	redisClient, err := redis.NewRedisClient(ctx, cfg.RedisURL)
	if err != nil {
		logrus.Fatal("Failed to connect to redis")
	}

	ethClient, err := ethclient.Dial(cfg.EthereumRpcURL)
	if err != nil {
		logrus.Fatal("Failed to connect to ethereum network")
	}

	service := watcher.NewWatcherService(logger, &cfg, qi, redisClient, ethClient)
	service.Run(ctx)
}
