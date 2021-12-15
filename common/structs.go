package common

import (
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/ethereum/go-ethereum/common"
)

const (
	AuthenticationHeaderKey = "Authentication"
	AuthenticationMessage   = "Authentication"
)

const (
	SignatureTypeNoPrepend   = 0
	SignatureTypeDecimal     = 1
	SignatureTypeHexadecimal = 2
)

// flag for contract method trade()
const (
	TraderFlagIsNullByte             = 0
	TraderFlagIsBuyByte              = 1
	TraderFlagIsSellByte             = 2
	TraderFlagIsNegativeLimitFeeByte = 4
)

// redis keys
const (
	WebsocketMessageQueueKey = "websocket_messages_queue"
	EngineEventQueueKey      = "engine_events_queue"
	OrderBookSnapshotKey     = "order_book_snapshot:"
)

// engine event
const (
	EventNewOrder           = "engine_event_new_order"
	EventCancelOrder        = "engine_event_cancel_order"
	EventConfirmTransaction = "engine_event_confirm_transaction"
)

type Event struct {
	Type     string `json:"event_type"`
	MarketID string `json:"market_id"`
}

type NewOrderEvent struct {
	Type     string       `json:"event_type"`
	MarketID string       `json:"market_id"`
	Order    models.Order `json:"order"`
}

type CancelOrderEvent struct {
	Type     string `json:"event_type"`
	MarketID string `json:"market_id"`
	ID       string `json:"id"`
}

type ConfirmTransactionEvent struct {
	Type      string `json:"event_type"`
	MarketID  string `json:"market_id"`
	Hash      string `json:"hash"`
	Status    string `json:"status"`
	Timestamp uint64 `json:"timestamp"`
}

// TransactionData common struct for engine and launcher
type TransactionData struct {
	Accounts []common.Address             `json:"accounts"`
	Trades   []contract.DxlnTradeTradeArg `json:"trades"`
}

// Websocket message

const (
	WsUpdateOrderBookPriceType = "update_order_book_price"
	WsAccountNewOrderType      = "account_new_order"
	WsAccountUpdateOrderType   = "account_update_order"
	WsMarketNewTradeType       = "market_new_trade"
	WsAccountNewTradeType      = "account_new_trade"
	WsAccountUpdateTradeType   = "account_update_trade"
)

type WebSocketMessage struct {
	ChannelID string               `json:"channel_id"`
	Data      WebSocketMessageData `json:"data"`
}

type WebSocketMessageData struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type WebsocketUpdateOrderBookPricePayload struct {
	Sequence uint64 `json:"sequence"`
	Side     string `json:"side"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type WebsocketMarketNewTradePayload struct {
	Amount     string    `json:"amount"`
	Price      string    `json:"price"`
	ExecutedAt time.Time `json:"executed_at"`
}

type WebsocketAccountOrderPayload struct {
	Order *models.Order `json:"order"`
}

type WebsocketAccountTradePayload struct {
	Trade *models.Trade `json:"trade"`
}
