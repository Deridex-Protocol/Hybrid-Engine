package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MarketHandler struct {
	qi    models.QI
	log   *logrus.Entry
	redis *redis.Client
}

func NewMarketHandler(qi models.QI, log *logrus.Entry, redis *redis.Client) MarketHandler {
	return MarketHandler{
		qi:    qi,
		log:   log,
		redis: redis,
	}
}

// GetOrderBook returns order book of the market
// @Summary Get order book of the market by market id
// @Description Returns market order book. Response contains universal data structure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty
// @Tags Markets
// @Produce json
// @Param market_id path int true "Market ID"
// @Success 200 {object} api.Response{data=handlers.Snapshot} "Market orderbook is inside data field"
// @Router /markets/{market_id}/orderbook [get]
func (h MarketHandler) GetOrderBook(p interface{}) (interface{}, error) {
	params := p.(*OrderBookReq)
	marketID := params.MarketID

	orderBookStr, err := h.redis.Get(common.OrderBookSnapshotKey + marketID).Result()
	if err != nil && err != redis.Nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err != nil {
		orderBookJSON, err := json.Marshal(&Snapshot{
			Sequence: 0,
			Bids:     [][2]string{},
			Asks:     [][2]string{},
		})
		if err != nil {
			h.log.WithError(err).Error("Failed to marshal snapshot struct")
			return nil, echo.NewHTTPError(http.StatusInternalServerError)
		}
		orderBookStr = string(orderBookJSON)
	}

	var snapshot Snapshot
	if err = json.Unmarshal([]byte(orderBookStr), &snapshot); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	if params.PerPage > 0 {
		if len(snapshot.Asks) > params.PerPage {
			snapshot.Asks = snapshot.Asks[:params.PerPage]
		}
		if len(snapshot.Bids) > params.PerPage {
			snapshot.Bids = snapshot.Bids[:params.PerPage]
		}
	}

	return snapshot, nil
}

// GetMarkets returns list of existing markets
// @Summary Get list of markets
// @Description Returns list of market. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty
// @Tags Markets
// @Produce json
// @Success 200 {object} api.Response{data=[]handlers.Market} "Array of existing markets"
// @Router /markets [get]
func (h MarketHandler) GetMarkets(_ interface{}) (interface{}, error) {
	var markets []Market

	dbMarkets, err := h.qi.Market().FindPublishedMarkets()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	yesterday := time.Now().UTC().Add(-time.Hour * 24)
	for _, dbMarket := range dbMarkets {
		var marketStatus MarketStatus

		marketStatusDB, err := h.qi.Trade().GetMarketStatus(dbMarket.ID, yesterday, time.Now())
		if err != nil {
			h.log.WithError(err).Error("Failed to get market status")
			marketStatus = MarketStatus{
				LastPrice:           decimal.Zero,
				Price24h:            decimal.Zero,
				Amount24h:           decimal.Zero,
				QuoteTokenVolume24h: decimal.Zero,
			}
		} else {
			marketStatus = MarketStatus{
				LastPrice:           marketStatusDB.LastPrice,
				Price24h:            marketStatusDB.LastPrice.Sub(marketStatusDB.FirstPrice).Div(marketStatusDB.FirstPrice),
				Amount24h:           marketStatusDB.Amount24h,
				QuoteTokenVolume24h: marketStatusDB.QuoteTokenVolume24h,
			}
		}

		markets = append(markets, Market{
			ID:                  dbMarket.ID,
			BaseToken:           dbMarket.BaseTokenSymbol,
			BaseTokenName:       dbMarket.BaseTokenName,
			BaseTokenDecimals:   dbMarket.BaseTokenDecimals,
			BaseTokenAddress:    dbMarket.BaseTokenAddress,
			QuoteToken:          dbMarket.QuoteTokenSymbol,
			QuoteTokenName:      dbMarket.QuoteTokenName,
			QuoteTokenDecimals:  dbMarket.QuoteTokenDecimals,
			QuoteTokenAddress:   dbMarket.QuoteTokenAddress,
			MinOrderSize:        dbMarket.MinOrderSize,
			PriceDecimals:       dbMarket.PriceDecimals,
			AmountDecimals:      dbMarket.AmountDecimals,
			SupportedOrderTypes: []string{"limit", "market"},
			MarketStatus:        marketStatus,
		})
	}

	return markets, nil
}
