package bot

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/crypto"
	"bitbucket.ideasoft.io/dex/dex-backend/common/yahoo"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Service struct {
	log         *logrus.Entry
	cfg         *Config
	yahooClient yahoo.YahooClient
}

func NewMakerService(log *logrus.Entry, cfg *Config) Service {
	return Service{
		log:         log,
		cfg:         cfg,
		yahooClient: yahoo.NewYahooClient(),
	}
}

func (s *Service) Run(ctx context.Context) {
	var err error
	var isFirstAddress bool

	rand.Seed(time.Now().UnixNano())

	price, err := s.yahooClient.GetYahooPrice("TSLA")
	if err != nil {
		s.log.WithError(err).Error("Failed get price from yahoo")
		return
	}

	updateYahooPriceTicker := time.NewTicker(5 * time.Minute)
	botTickerTime, err := time.ParseDuration(s.cfg.TickerTime)
	if err != nil {
		s.log.WithError(err).Error("Failed to parse bot ticker time")
		return
	}
	botTicker := time.NewTicker(botTickerTime)

	for {
		select {
		case <-updateYahooPriceTicker.C:
			price, err = s.yahooClient.GetYahooPrice("TSLA")
			if err != nil {
				s.log.WithError(err).Error("Failed get price from yahoo")
				continue
			}
			if price.IsZero() {
				s.log.WithError(err).Error("Price from yahoo is zero")
				continue
			}
			s.log.WithField("price", price).Info("Successfully get yahoo price")
		case <-botTicker.C:
			amount := fmt.Sprintf("%f", randomNumber(s.cfg.MinAmount, s.cfg.MaxAmount, 6))

			minPrice, _ := price.Mul(decimal.New(9, -1)).Float64()
			maxPrice, _ := price.Mul(decimal.New(11, -1)).Float64()
			orderPrice := fmt.Sprintf("%f", randomNumber(minPrice, maxPrice, 6))

			if isFirstAddress {
				s.placeOrder(amount, orderPrice, "buy", s.cfg.PrivateKey1, s.cfg.Address1)
			} else {
				s.placeOrder(amount, orderPrice, "sell", s.cfg.PrivateKey2, s.cfg.Address2)
			}

			isFirstAddress = !isFirstAddress

		case <-ctx.Done():
			return
		}
	}
}

func (s *Service) placeOrder(amount, price, side, privateKey, address string) {
	body, err := json.Marshal(map[string]interface{}{
		"amount":     amount,
		"price":      price,
		"side":       side,
		"order_type": "limit",
		"market_id":  s.cfg.MarketID,
	})
	if err != nil {
		s.log.WithError(err).Error("Failed to marshal body order")
		return
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/orders/build/limit", s.cfg.ApiURL), bytes.NewReader(body))
	if err != nil {
		s.log.WithError(err).Error("Failed to create request")
		return
	}
	req.Header.Add("Content-Type", "application/json")

	authMessage := common.AuthenticationMessage + ":" + time.Now().AddDate(1, 0, 0).Format(time.RFC3339)

	authHeader, err := crypto.SignAuthHeader(privateKey, address, authMessage)
	if err != nil {
		s.log.WithError(err).Error("Failed to sign auth header")
		return
	}
	req.Header.Add(common.AuthenticationHeaderKey, authHeader)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		s.log.WithError(err).Error("Failed to build order")
		return
	}

	s.log.WithField("order_body", string(body)).Info("Build order success")

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.log.WithError(err).Errorf("Failed to read response body")
		return
	}
	if err = res.Body.Close(); err != nil {
		s.log.WithError(err).Errorf("Failed to close response body")
		return
	}

	var buildOrderRes struct {
		Data struct {
			OrderID string `json:"order_id"`
		} `json:"data"`
	}

	if err = json.Unmarshal(resBytes, &buildOrderRes); err != nil {
		s.log.WithError(err).Errorf("Failed to unmasrshank build order response")
		return
	}

	if buildOrderRes.Data.OrderID == "" {
		s.log.Error("Failed to build order. Empty order id")
		return
	}

	orderHash, err := hex.DecodeString(buildOrderRes.Data.OrderID[2:])
	if err != nil {
		s.log.WithError(err).Errorf("Failed to decode order id to bytes")
		return
	}

	signature, err := crypto.SignEthereumMessage(privateKey, orderHash)
	if err != nil {
		s.log.WithError(err).
			WithField("order_id", buildOrderRes.Data.OrderID).
			Error("Failed to create signature for order id")
		return
	}

	placeOrderRequestBody, err := json.Marshal(map[string]interface{}{
		"order_id":  buildOrderRes.Data.OrderID,
		"signature": "0x" + hex.EncodeToString(signature),
	})
	if err != nil {
		s.log.WithError(err).Error("failed to marshal body: %v", err)
		return
	}

	req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/orders", s.cfg.ApiURL), bytes.NewReader(placeOrderRequestBody))
	if err != nil {
		s.log.WithError(err).Error("Failed to create request")
		return
	}
	req.Header.Add("Content-Type", "application/json")

	authHeader, err = crypto.SignAuthHeader(privateKey, address, authMessage)
	if err != nil {
		s.log.WithError(err).Error("Failed to sign auth header")
		return
	}
	req.Header.Add(common.AuthenticationHeaderKey, authHeader)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		s.log.WithError(err).Error("Failed to place order")
		return
	}

	s.log.WithField("order_id", buildOrderRes.Data.OrderID).Info("Place order success")
}

func randomNumber(min, max, decimals float64) float64 {
	r := rand.Float64()*(max-min) + min
	pow := math.Pow(10, decimals)
	return math.Floor(r*pow) / pow
}
