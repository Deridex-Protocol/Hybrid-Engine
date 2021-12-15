package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

type QI interface {
	Market() MarketQI
	Order() OrderQI
	Token() TokenQI
	Trade() TradeQI
	Transaction() TransactionQI
}

type qi struct {
	db *gorm.DB
}

func (q qi) Market() MarketQI {
	return marketQI{db: q.db}
}

func (q qi) Order() OrderQI {
	return orderQI{db: q.db}
}

func (q qi) Token() TokenQI {
	return tokenQI{db: q.db}
}

func (q qi) Trade() TradeQI {
	return tradeQI{db: q.db}
}

func (q qi) Transaction() TransactionQI {
	return transactionQI{db: q.db}
}

func Connect(url string) (QI, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}

	return &qi{db: db}, nil
}
