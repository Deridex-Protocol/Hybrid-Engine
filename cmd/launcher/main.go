package main

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"bitbucket.ideasoft.io/dex/dex-backend/common/shutdown"
	"bitbucket.ideasoft.io/dex/dex-backend/services/launcher"
	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := shutdown.GracefulShutdown()

	var cfg launcher.Config
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatal("Failed to parse envs to struct")
	}

	log := logrus.New()
	logger := log.WithField("service", "launcher")

	qi, err := models.Connect(cfg.DatabaseURL)
	if err != nil {
		logrus.Fatal("Failed to connect to database")
	}

	ethClient, err := ethclient.Dial(cfg.EthereumRpcURL)
	if err != nil {
		logrus.Fatal("Failed to connect to ethereum network")
	}

	contractClient, err := contract.NewContract(common.HexToAddress(cfg.ProxyAddress), ethClient)
	if err != nil {
		logrus.Fatal("Failed to create contract")
	}

	service := launcher.NewLauncherService(logger, &cfg, qi, ethClient, contractClient)
	service.Run(ctx)
}
