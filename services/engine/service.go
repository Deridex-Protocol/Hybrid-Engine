package engine

import (
	"context"
	"encoding/json"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log       *logrus.Entry
	cfg       *Config
	qi        models.QI
	redis     *redis.Client
	ethClient *ethclient.Client
	markets   map[string]*MarketHandler
}

func NewEngineService(log *logrus.Entry, cfg *Config, qi models.QI, redis *redis.Client) Service {
	return Service{
		log:     log,
		cfg:     cfg,
		qi:      qi,
		redis:   redis,
		markets: make(map[string]*MarketHandler),
	}
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Engine service start")

	if err := s.initMarkets(); err != nil {
		s.log.WithError(err).Error("Failed to init markets")
		return
	}

	for i := range s.markets {
		go s.runMarket(ctx, s.markets[i])
	}

	for {
		select {
		case <-ctx.Done():
			s.log.Info("Engine service stopped")
			return
		default:
			// listen redis engine event queue
			res, err := s.redis.BRPop(time.Second, common.EngineEventQueueKey).Result()
			if err != nil && err != redis.Nil {
				s.log.WithError(err).Error("Failed pop message from event queue")
				return
			}
			if err == redis.Nil {
				continue
			}

			data := []byte(res[1])

			var event common.Event
			if err := json.Unmarshal(data, &event); err != nil {
				s.log.WithError(err).Error("Failed pop message from event queue")
				continue
			}

			marketHandler, ok := s.markets[event.MarketID]
			if !ok {
				s.log.WithField("market_id", event.MarketID).Error("Market not found in engine")
				continue
			}

			marketHandler.eventChan <- data
		}
	}
}

func (s *Service) initMarkets() error {
	markets, err := s.qi.Market().FindPublishedMarkets()
	if err != nil {
		return err
	}

	for i := range markets {
		if err := s.initMarket(markets[i]); err != nil {
			return err
		}
		s.log.WithField("market_id", markets[i].ID).Info("Market init done")
	}

	return nil
}

func (s *Service) initMarket(market *models.Market) error {
	if _, ok := s.markets[market.ID]; ok {
		s.log.WithField("market_id", market.ID).Info("Open market is failed. Market already exist")
		return nil
	}

	if !market.IsPublished {
		s.log.WithField("market_id", market.ID).Info("Open market is failed. Market not published")
		return nil
	}

	marketHandler, err := NewMarketHandler(s.log, s.qi, s.redis, market, s.cfg.ExchangeAddress, s.cfg.RelayerAddress)
	if err != nil {
		return err
	}

	s.markets[market.ID] = marketHandler

	return nil
}

func (s *Service) runMarket(ctx context.Context, marketHandler *MarketHandler) {
	s.log.WithField("market_id", marketHandler.market.ID).Info("Market Handler is running")
	marketHandler.Run(ctx)
	s.log.WithField("market_id", marketHandler.market.ID).Info("Market Handler is stopped")
}
