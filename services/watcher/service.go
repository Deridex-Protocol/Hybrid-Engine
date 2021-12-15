package watcher

import (
	"context"
	"database/sql"
	"encoding/json"
	"math/big"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const sleepSecondsForNewBlock = 10 * time.Second

type Service struct {
	log       *logrus.Entry
	cfg       *Config
	qi        models.QI
	redis     *redis.Client
	ethClient *ethclient.Client
}

func NewWatcherService(log *logrus.Entry, cfg *Config, qi models.QI, redis *redis.Client, ethClient *ethclient.Client) Service {
	return Service{
		log:       log,
		cfg:       cfg,
		qi:        qi,
		redis:     redis,
		ethClient: ethClient,
	}
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Watcher service start")

	latestBlockNumber, err := s.ethClient.BlockNumber(ctx)
	if err != nil {
		s.log.WithError(err).Error("Failed get block number from redis")
		return
	}
	blockNumber := big.NewInt(int64(latestBlockNumber) + 1)

	transactions, err := s.qi.Transaction().FindAllPending()
	if err != nil {
		s.log.WithError(err).Error("Failed get all created transactions")
		return
	}
	for i := range transactions {
		if !transactions[i].Hash.Valid {
			continue
		}

		tx, isPending, err := s.ethClient.TransactionByHash(ctx, ethcommon.HexToHash(transactions[i].Hash.String))
		if err != nil {
			s.log.WithError(err).
				WithField("tx_hash", transactions[i].Hash.String).
				Error("Failed get transaction by hash from ethereum")
			return
		}
		if isPending {
			continue
		}

		if err := s.handlerResultTransaction(ctx, tx, nil); err != nil {
			s.log.WithError(err).
				WithField("tx_hash", transactions[i].Hash.String).
				Error("Failed get handler transaction result")
			return
		}
	}

	for {
		select {
		case <-ctx.Done():
			s.log.Info("Watcher service stop")
			return
		default:
			block, err := s.ethClient.BlockByNumber(ctx, blockNumber)
			if err != nil && err != ethereum.NotFound {
				s.log.WithError(err).Error("Failed to get block by number")
				continue
			}
			if err == ethereum.NotFound {
				time.Sleep(sleepSecondsForNewBlock)
				continue
			}

			for _, tx := range block.Transactions() {
				if err := s.handlerResultTransaction(ctx, tx, block); err != nil {
					s.log.WithError(err).
						WithField("tx_hash", tx.Hash().String).
						Error("Failed get handler transaction result")
					return
				}
			}

			blockNumber.Add(blockNumber, big.NewInt(1))
		}
	}
}

func (s *Service) handlerResultTransaction(ctx context.Context, tx *types.Transaction, block *types.Block) error {
	transaction, err := s.qi.Transaction().FindByHash(tx.Hash().Hex())
	if err != nil && err != gorm.ErrRecordNotFound {
		s.log.WithError(err).
			WithField("tx_hash", tx.Hash().Hex()).
			Error("Failed get transaction by hash from db")
		return err
	}
	if err == gorm.ErrRecordNotFound || !transaction.Hash.Valid || transaction.Status != models.TradeStatusPending {
		return nil
	}

	receipt, err := s.ethClient.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		s.log.WithError(err).
			WithField("tx_hash", tx.Hash().Hex()).
			Error("Failed to get tx receipt status")
		return err
	}

	if block == nil {
		block, err = s.ethClient.BlockByHash(ctx, receipt.BlockHash)
		if err != nil {
			s.log.WithError(err).Error("Failed to get block by hash")
			return err
		}
	}

	s.log.WithField("tx_hash", tx.Hash().Hex()).
		WithField("tx_result", receipt.Status).
		Info("Transaction result for existing hash")

	if receipt.Status == types.ReceiptStatusSuccessful {
		transaction.Status = models.TransactionStatusSuccessful
	} else {
		transaction.Status = models.TransactionStatusFailed
	}
	transaction.BlockNumber = sql.NullInt64{Int64: receipt.BlockNumber.Int64(), Valid: true}
	transaction.GasUsed = sql.NullInt64{Int64: int64(receipt.GasUsed), Valid: true}

	if err := s.qi.Transaction().UpdateTransaction(transaction); err != nil {
		s.log.WithError(err).Error("Failed to update transaction")
		return err
	}

	event := &common.ConfirmTransactionEvent{
		Type:      common.EventConfirmTransaction,
		MarketID:  transaction.MarketID,
		Hash:      tx.Hash().Hex(),
		Status:    transaction.Status,
		Timestamp: block.Time(),
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		s.log.WithError(err).Error("Failed to marshal event")
		return err
	}

	if err := s.redis.LPush(common.EngineEventQueueKey, eventJSON).Err(); err != nil {
		s.log.WithError(err).Errorf("Failed to push engine event into redis queue")
	}

	return nil
}
