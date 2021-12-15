package main

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/shutdown"
	"bitbucket.ideasoft.io/dex/dex-backend/services/bot"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := shutdown.GracefulShutdown()

	var cfg bot.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal("Failed to parse envs to struct")
	}

	log := logrus.New()
	logger := log.WithField("service", "bot")

	service := bot.NewMakerService(logger, &cfg)
	service.Run(ctx)
}
