package handlers

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/shopspring/decimal"
)

type BaseReq struct {
	Address string `json:"address"`
}

func (b *BaseReq) GetAddress() string {
	return b.Address
}

func (b *BaseReq) SetAddress(address string) {
	b.Address = address
}

// Requests
type (
	OrderBookReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id" param:"market_id" validate:"required"`
		PerPage  int    `json:"per_page"  query:"per_page"`
	}

	CandlesReq struct {
		BaseReq     `swaggerignore:"true"`
		MarketID    string `json:"market_id"    param:"market_id"    validate:"required"`
		From        int64  `json:"from"        query:"from"        validate:"required"`
		To          int64  `json:"to"          query:"to"          validate:"required"`
		Granularity int64  `json:"granularity" query:"granularity" validate:"required"`
	}

	QueryOrderReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id" query:"market_id" validate:"required"`
		Status   string `json:"status"    query:"status" validate:"max=128"`
		Page     int    `json:"page"      query:"page"`
		PerPage  int    `json:"per_page"  query:"per_page"`
	}

	GetOrdersInfoReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id" query:"market_id" validate:"required"`
	}

	GetOrdersInfoResp struct {
		UnrealizedPNL decimal.Decimal `json:"unrealized_pnl"`
	}

	QueryTradeReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id" param:"market_id" validate:"required"`
		Status   string `json:"status"    query:"status"`
		Page     int    `json:"page"      query:"page"`
		PerPage  int    `json:"per_page"  query:"per_page"`
	}

	QuerySingleOrderReq struct {
		BaseReq `swaggerignore:"true"`
		OrderID string `json:"order_id" param:"order_id" validate:"required"`
	}

	BuildLimitOrderReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id"  validate:"required"`
		Side     string `json:"side"       validate:"required,oneof=buy sell"`
		Price    string `json:"price"      validate:"required"`
		Amount   string `json:"amount"     validate:"required"`
	}

	BuildMarketOrderReq struct {
		BaseReq  `swaggerignore:"true"`
		MarketID string `json:"market_id" validate:"required"`
		Side     string `json:"side"      validate:"required,oneof=buy sell"`
		Amount   string `json:"amount"    validate:"required"`
	}

	PlaceOrderReq struct {
		BaseReq   `swaggerignore:"true"`
		OrderID   string `json:"order_id"  validate:"required,len=66"`
		Signature string `json:"signature" validate:"required"`
	}

	CancelOrderReq struct {
		BaseReq `swaggerignore:"true"`
		ID      string `json:"id" param:"order_id" validate:"required,len=66"`
	}

	BuildOrderResp struct {
		OrderID string `json:"order_id"`
	}

	OrderResp struct {
		Order *models.Order `json:"order"`
	}

	OrdersResp struct {
		Count  int64           `json:"count"`
		Orders []*models.Order `json:"orders"`
	}

	TradesResp struct {
		Count  int64           `json:"count"`
		Trades []*models.Trade `json:"trades"`
	}

	CandlesResp struct {
		Candles []*Bar `json:"candles"`
	}
)

// Models
type (
	Market struct {
		ID                  string          `json:"id"`
		BaseToken           string          `json:"base_token"`
		BaseTokenName       string          `json:"base_token_name"`
		BaseTokenDecimals   int             `json:"base_token_decimals"`
		BaseTokenAddress    string          `json:"base_token_address"`
		QuoteToken          string          `json:"quote_token"`
		QuoteTokenName      string          `json:"quote_token_name"`
		QuoteTokenDecimals  int             `json:"quote_token_decimals"`
		QuoteTokenAddress   string          `json:"quote_token_address"`
		MinOrderSize        decimal.Decimal `json:"min_order_size"`
		PriceDecimals       int             `json:"price_decimals"`
		AmountDecimals      int             `json:"amount_decimals"`
		SupportedOrderTypes []string        `json:"supported_order_types"`
		MarketStatus
	}

	MarketStatus struct {
		LastPrice           decimal.Decimal `json:"last_price"`
		Price24h            decimal.Decimal `json:"price24h"`
		Amount24h           decimal.Decimal `json:"amount24h"`
		QuoteTokenVolume24h decimal.Decimal `json:"quote_token_volume24h"`
	}

	Bar struct {
		Time   int64           `json:"time"`
		Open   decimal.Decimal `json:"open"`
		Close  decimal.Decimal `json:"close"`
		Low    decimal.Decimal `json:"low"`
		High   decimal.Decimal `json:"high"`
		Volume decimal.Decimal `json:"volume"`
	}

	CacheOrder struct {
		ID       string          `json:"id"`
		MarketID string          `json:"market_id"`
		Side     string          `json:"side"`
		Type     string          `json:"type"`
		Price    decimal.Decimal `json:"price"`
		Amount   decimal.Decimal `json:"amount"`
		Address  string          `json:"address"`
		Flags    [32]byte        `json:"flags"`
	}

	Snapshot struct {
		Sequence uint64      `json:"sequence"`
		Bids     [][2]string `json:"bids"`
		Asks     [][2]string `json:"asks"`
	}

	BuildOrder struct {
		MarketID  string
		Side      string
		Price     string
		Amount    string
		OrderType string
	}
)
