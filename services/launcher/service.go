package launcher

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"encoding/json"
	"math/big"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log            *logrus.Entry
	cfg            *Config
	qi             models.QI
	ethClient      *ethclient.Client
	contractClient *contract.Contract
	privateKey     *ecdsa.PrivateKey
	nonce          *big.Int
}

func NewLauncherService(log *logrus.Entry, cfg *Config, qi models.QI, ethClient *ethclient.Client,
	contractClient *contract.Contract) Service {
	return Service{
		log:            log,
		cfg:            cfg,
		qi:             qi,
		ethClient:      ethClient,
		contractClient: contractClient,
	}
}

func (s *Service) Run(ctx context.Context) {
	s.log.Info("Launcher service start")

	privateKey, err := crypto.HexToECDSA(s.cfg.RelayerPrivateKey)
	if err != nil {
		s.log.WithError(err).Error("Failed to parse private key")
		return
	}
	s.privateKey = privateKey

	if err = s.UpdateNonce(ctx); err != nil {
		s.log.WithError(err).Error("Failed to update nonce")
		return
	}

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			s.handleTick(ctx)
		case <-ctx.Done():
			s.log.Info("Launcher service stop")
			return
		}
	}
}

func (s *Service) handleTick(ctx context.Context) {
	transactions, err := s.qi.Transaction().FindAllCreated()
	if err != nil {
		s.log.WithError(err).Error("Failed to get all created transaction")
		return
	}

	for _, modelTransaction := range transactions {
		var transactionData common.TransactionData
		if err = json.Unmarshal([]byte(modelTransaction.Data), &transactionData); err != nil {
			s.log.WithError(err).Error("Failed to unmarshal transaction data")
			continue
		}

		gasPrice := decimal.New(s.cfg.ManualGasPrice, 0)
		if s.cfg.ManualGasPrice == 0 {
			gasPrice, err = GasPriceInWei()
			if err != nil {
				s.log.WithError(err).Error("Failed to get gas price")
				gasPrice = decimal.New(3, 9) // 3Gwei
			}
		}

		opts := &bind.TransactOpts{
			From:  crypto.PubkeyToAddress(s.privateKey.PublicKey),
			Nonce: s.nonce,
			Signer: func(_ ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
				signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), s.privateKey)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(types.HomesteadSigner{}, signature)
			},
			GasPrice: gasPrice.BigInt(),
			GasLimit: s.cfg.GasLimit,
		}

		tx, err := s.contractClient.Trade(opts, transactionData.Accounts, transactionData.Trades)
		if err != nil {
			s.log.WithError(err).
				WithField("transaction_id", modelTransaction.ID).
				Info("Send Tx failed")

			modelTransaction.Nonce = sql.NullInt64{Int64: s.nonce.Int64(), Valid: true}
			modelTransaction.Status = models.TransactionStatusFailed

			if err = s.UpdateTransaction(modelTransaction); err != nil {
				s.log.WithError(err).
					WithField("transaction_id", modelTransaction.ID).
					Error("Failed to update transaction")
			}

			if err = s.UpdateNonce(ctx); err != nil {
				s.log.WithError(err).Error("Failed to update nonce")
				return
			}
			continue
		}

		s.log.WithField("transaction_id", modelTransaction.ID).
			WithField("tx_hash", tx.Hash().String()).
			Info("Send Tx")

		modelTransaction.Status = models.TransactionStatusPending
		modelTransaction.Hash = sql.NullString{Valid: true, String: tx.Hash().String()}
		modelTransaction.GasLimit = sql.NullInt64{Int64: int64(opts.GasLimit), Valid: true}
		modelTransaction.GasPrice = decimal.NullDecimal{Decimal: gasPrice, Valid: true}
		modelTransaction.Nonce = sql.NullInt64{Int64: s.nonce.Int64(), Valid: true}

		s.nonce = s.nonce.Add(s.nonce, big.NewInt(1))

		if err = s.UpdateTransaction(modelTransaction); err != nil {
			s.log.WithError(err).
				WithField("transaction_id", modelTransaction.ID).
				Error("Failed to update transaction hash")
			continue
		}
	}
}

func (s *Service) UpdateTransaction(transaction *models.Transaction) (err error) {
	if err = s.qi.Transaction().UpdateTransaction(transaction); err != nil {
		s.log.WithError(err).Error("Failed to update transaction")
		return
	}

	trades, err := s.qi.Trade().FindTradeByTransactionID(transaction.ID)
	if err != nil {
		s.log.WithError(err).Error("Failed to get trade by tx hash")
		return
	}

	for _, trade := range trades {
		trade.TransactionHash = transaction.Hash.String
		if err = s.qi.Trade().UpdateTrade(trade); err != nil {
			s.log.WithError(err).Error("Failed to update trade")
			return
		}
	}
	return
}

func (s *Service) UpdateNonce(ctx context.Context) error {
	addressNonce, err := s.ethClient.NonceAt(ctx, crypto.PubkeyToAddress(s.privateKey.PublicKey), nil)
	if err != nil {
		return err
	}
	s.nonce = big.NewInt(int64(addressNonce))
	s.log.WithField("nonce", s.nonce).Info("Update nonce successfully")
	return nil
}
