package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

const (
	TradeStatusPending    = "pending"
	TradeStatusSuccessful = "successful"
	TradeStatusFailed     = "failed"
)

type TradeQI interface {
	FindAllTrades(marketID string, limit, offset int) (int64, []*Trade, error)
	FindTradesByMarket(marketID string, startTime time.Time, endTime time.Time) ([]*Trade, error)
	FindTradeByTransactionID(transactionID int64) ([]*Trade, error)
	FindTradesByHash(hash string) ([]*Trade, error)
	FindTradeByID(id int64) (*Trade, error)
	FindAccountMarketTrades(account, marketID, status string, limit, offset int) (int64, []*Trade, error)
	GetMarketStatus(marketID string, startTime time.Time, endTime time.Time) (*MarketStatus, error)
	InsertTrade(trade *Trade) error
	UpdateTrade(trade *Trade) error
}

type Trade struct {
	ID              int64           `json:"id"               db:"id" primaryKey:"true" autoIncrement:"true" gorm:"primary_key"`
	TransactionID   int64           `json:"transaction_id"   db:"transaction_id"`
	TransactionHash string          `json:"transaction_hash" db:"transaction_hash"`
	Status          string          `json:"status"           db:"status"`
	MarketID        string          `json:"market_id"        db:"market_id"`
	Maker           string          `json:"maker"            db:"maker"`
	Taker           string          `json:"taker"            db:"taker"`
	TakerSide       string          `json:"taker_side"       db:"taker_side"`
	MakerOrderID    string          `json:"maker_order_id"   db:"maker_order_id"`
	TakerOrderID    string          `json:"taker_order_id"   db:"taker_order_id"`
	Amount          decimal.Decimal `json:"amount"           db:"amount"`
	Price           decimal.Decimal `json:"price"            db:"price"`
	ExecutedAt      time.Time       `json:"executed_at"      db:"executed_at"`
	CreatedAt       time.Time       `json:"created_at"       db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"       db:"updated_at"`
}

func (Trade) TableName() string {
	return "trades"
}

type tradeQI struct {
	db *gorm.DB
}

func (q tradeQI) FindAllTrades(marketID string, limit, offset int) (int64, []*Trade, error) {
	var trades []*Trade
	var count int64
	err := q.db.Where("market_id = ? and status = ?", marketID, TradeStatusSuccessful).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&trades).
		Count(&count).Error
	return count, trades, err
}

func (q tradeQI) FindTradesByMarket(marketID string, startTime time.Time, endTime time.Time) ([]*Trade, error) {
	var trades []*Trade
	err := q.db.Where("market_id = ? and status = ? and executed_at between ? and ? ", marketID, TradeStatusSuccessful, startTime, endTime).
		Order("executed_at desc").Find(&trades).Error
	return trades, err
}

func (q tradeQI) FindTradeByTransactionID(transactionID int64) ([]*Trade, error) {
	var trades []*Trade
	err := q.db.Where("transaction_id = ? ", transactionID).Order("created_at asc").Find(&trades).Error
	return trades, err
}

func (q tradeQI) FindTradesByHash(hash string) ([]*Trade, error) {
	var trades []*Trade
	err := q.db.Where("transaction_hash = ?", hash).Order("created_at desc").Find(&trades).Error
	return trades, err
}

func (q tradeQI) FindTradeByID(id int64) (*Trade, error) {
	var trade Trade
	err := q.db.Where("id = ?", id).Find(&trade).Error
	return &trade, err
}

func (q tradeQI) FindAccountMarketTrades(account, marketID, status string, limit, offset int) (int64, []*Trade, error) {
	var trades []*Trade
	var count int64

	err := q.db.Where("market_id = ? and (taker = ? or maker = ?)", marketID, account, account).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&trades).
		Count(&count).Error

	return count, trades, err
}

type MarketStatus struct {
	FirstPrice          decimal.Decimal `db:"first_price"`
	LastPrice           decimal.Decimal `db:"last_price"`
	QuoteTokenVolume24h decimal.Decimal `db:"quote_token_volume24h"`
	Amount24h           decimal.Decimal `db:"amount24h"`
}

func (q tradeQI) GetMarketStatus(marketID string, startTime time.Time, endTime time.Time) (*MarketStatus, error) {
	row := q.db.Raw(`
		select min(price) as first_price,
			max(price) as last_price,
			sum(amount) as quote_token_volume24h,
			sum(amount*price) as amount24h
		from trades
		where market_id = $1 and status = $2 and executed_at between $3 and $4`,
		marketID, TradeStatusSuccessful, startTime, endTime).Row()
	if row == nil {
		return nil, nil
	}

	var marketStatus MarketStatus
	if err := row.Scan(&marketStatus); err != nil {
		return nil, err
	}
	return &marketStatus, nil
}

func (q tradeQI) InsertTrade(trade *Trade) error {
	return q.db.Create(trade).Error
}

func (q tradeQI) UpdateTrade(trade *Trade) error {
	return q.db.Save(trade).Error
}
