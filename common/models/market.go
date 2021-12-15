package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type MarketQI interface {
	FindPublishedMarkets() ([]*Market, error)
	FindMarketByID(marketID string) (*Market, error)
	InsertMarket(market *Market) error
	UpdateMarket(market *Market) error
}

type Market struct {
	ID                 string          `json:"id"                   db:"id" primaryKey:"true" gorm:"primary_key"`
	BaseTokenAddress   string          `json:"base_token_address"   db:"base_token_address"`
	BaseTokenSymbol    string          `json:"base_token_symbol"    db:"base_token_symbol"`
	BaseTokenName      string          `json:"base_token_name"      db:"base_token_name"`
	BaseTokenDecimals  int             `json:"base_token_decimals"  db:"base_token_decimals"`
	QuoteTokenAddress  string          `json:"quote_token_address"  db:"quote_token_address"`
	QuoteTokenSymbol   string          `json:"quote_token_symbol"   db:"quote_token_symbol"`
	QuoteTokenName     string          `json:"quote_token_name"     db:"quote_token_name"`
	QuoteTokenDecimals int             `json:"quote_token_decimals" db:"quote_token_decimals"`
	MinOrderSize       decimal.Decimal `json:"min_order_size"       db:"min_order_size"`
	PriceDecimals      int             `json:"price_decimals"       db:"price_decimals"`
	AmountDecimals     int             `json:"amount_decimals"      db:"amount_decimals"`
	MakerFeeRate       decimal.Decimal `json:"maker_fee_rate"       db:"maker_fee_rate"`
	TakerFeeRate       decimal.Decimal `json:"taker_fee_rate"       db:"taker_fee_rate"`
	IsPublished        bool            `json:"is_published"         db:"is_published"`
}

func (Market) TableName() string {
	return "markets"
}

type marketQI struct {
	db *gorm.DB
}

func (q marketQI) FindPublishedMarkets() ([]*Market, error) {
	var markets []*Market
	err := q.db.Where("is_published = ?", true).Find(&markets).Error
	return markets, err
}

func (q marketQI) FindMarketByID(marketID string) (*Market, error) {
	var market Market
	err := q.db.Where("id = ?", marketID).First(&market).Error
	return &market, err
}

func (q marketQI) InsertMarket(market *Market) error {
	return q.db.Create(market).Error
}

func (q marketQI) UpdateMarket(market *Market) error {
	return q.db.Save(market).Error
}
