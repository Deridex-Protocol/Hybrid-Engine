package handlers

import (
	"net/http"
	"sort"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type TradeHandler struct {
	qi  models.QI
	log *logrus.Entry
}

func NewTradeHandler(qi models.QI, log *logrus.Entry) TradeHandler {
	return TradeHandler{
		qi:  qi,
		log: log,
	}
}

const MaxBarsCount = 200

// GetAllTrades returns all trades of the market
// @Summary Get market trades by market id
// @Description Returns all market trades. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty
// @Tags Markets
// @Produce json
// @Param market_id path int true "Market ID"
// @Success 200 {object} api.Response{data=handlers.TradesResp} "Market trades are inside data field"
// @Router /markets/{market_id}/trades [get]
func (h *TradeHandler) GetAllTrades(p interface{}) (interface{}, error) {
	req := p.(*QueryTradeReq)
	if req.PerPage <= 0 {
		req.PerPage = 20
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := req.PerPage * (req.Page - 1)
	limit := req.PerPage

	count, trades, err := h.qi.Trade().FindAllTrades(req.MarketID, limit, offset)
	if err != nil {
		h.log.WithError(err).Error("Failed to get all trades")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return TradesResp{
		Count:  count,
		Trades: trades,
	}, nil
}

// GetAccountTrades returns account trades of the market
// @Summary Get account trades of the certain market by market id
// @Description Returns all account trades in the certain market. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Markets
// @Produce json
// @Param market_id path int true "Market ID"
// @Success 200 {object} api.Response{data=handlers.TradesResp} "Account trades are inside data field"
// @Router /markets/{market_id}/trades/mine [get]
func (h *TradeHandler) GetAccountTrades(p interface{}) (interface{}, error) {
	req := p.(*QueryTradeReq)
	if req.PerPage <= 0 {
		req.PerPage = 20
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	offset := req.PerPage * (req.Page - 1)
	limit := req.PerPage

	count, trades, err := h.qi.Trade().FindAccountMarketTrades(req.Address, req.MarketID, req.Status, limit, offset)
	if err != nil {
		h.log.WithError(err).Error("Failed to get account market trades")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return TradesResp{
		Count:  count,
		Trades: trades,
	}, nil
}

// GetTradingView returns trading view in form of candles
// @Summary Get trading view candles of the market by market id
// @Description Returns trading view candles of the market. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty
// @Tags Markets
// @Produce json
// @Param market_id path int true "Market ID"
// @Param from query int true "Period start"
// @Param to query int true "Period end"
// @Param granularity query int true "Data granularity in seconds"
// @Success 200 {object} api.Response{data=handlers.CandlesResp} "Candle values are inside data field"
// @Router /markets/{market_id}/candles [get]
func (h *TradeHandler) GetTradingView(p interface{}) (interface{}, error) {
	req := p.(*CandlesReq)
	from := req.From
	to := req.To
	granularity := req.Granularity

	if (to - granularity*MaxBarsCount) > from {
		from = to - granularity*MaxBarsCount
	}

	trades, err := h.qi.Trade().FindTradesByMarket(req.MarketID, time.Unix(from, 0), time.Unix(to, 0))
	if err != nil {
		h.log.WithError(err).Error("Failed to get market trades")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	if len(trades) == 0 {
		return map[string]interface{}{
			"candles": []*Bar{},
		}, nil
	}

	var candles []*Bar
	var currentIndex int64
	var currentBar *Bar

	sort.Slice(trades, func(i, j int) bool {
		return trades[i].ExecutedAt.Unix() < trades[j].ExecutedAt.Unix()
	})

	for _, trade := range trades {
		tIndex := trade.ExecutedAt.Unix() / granularity
		if currentBar == nil || currentBar.Volume.IsZero() {
			currentIndex = tIndex
			currentBar = &Bar{
				Time:   currentIndex * granularity,
				Volume: trade.Amount,
				Open:   trade.Price,
				Close:  trade.Price,
				High:   trade.Price,
				Low:    trade.Price,
			}
			continue
		}

		if tIndex < currentIndex+1 {
			currentBar.High = decimal.Max(currentBar.High, trade.Price)
			currentBar.Low = decimal.Min(currentBar.Low, trade.Price)
			currentBar.Volume = currentBar.Volume.Add(trade.Amount)
			currentBar.Close = trade.Price
		} else {
			currentIndex = tIndex
			if currentBar.Volume.IsZero() {
				continue
			}

			candles = append(candles, &Bar{
				Time:   currentBar.Time,
				Open:   currentBar.Open,
				Close:  currentBar.Close,
				Low:    currentBar.Low,
				High:   currentBar.High,
				Volume: currentBar.Volume,
			})

			currentBar = &Bar{
				Time:   currentIndex * granularity,
				Volume: trade.Amount,
				Open:   trade.Price,
				Close:  trade.Price,
				High:   trade.Price,
				Low:    trade.Price,
			}
		}
	}

	candles = append(candles, &Bar{
		Time:   currentBar.Time,
		Open:   currentBar.Open,
		Close:  currentBar.Close,
		Low:    currentBar.Low,
		High:   currentBar.High,
		Volume: currentBar.Volume,
	})

	return candles, nil
}
