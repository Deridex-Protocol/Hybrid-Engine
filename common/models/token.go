package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type TokenQI interface {
	GetAllTokens() ([]*Token, error)
	FindTokenBySymbol(string) (*Token, error)
	InsertToken(*Token) error
}

type Token struct {
	Symbol    string         `json:"symbol"     db:"symbol" gorm:"primary_key"`
	Name      string         `json:"name"       db:"name"`
	Address   sql.NullString `json:"address"    db:"address"`
	Decimals  int            `json:"decimals"   db:"decimals"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}

func (Token) TableName() string {
	return "tokens"
}

type tokenQI struct {
	db *gorm.DB
}

func (q tokenQI) GetAllTokens() ([]*Token, error) {
	var tokens []*Token
	err := q.db.Find(&tokens).Error
	return tokens, err
}

func (q tokenQI) FindTokenBySymbol(symbol string) (*Token, error) {
	var token Token
	err := q.db.Where("symbol = ?", symbol).Find(&token).Error
	return &token, err
}

func (q tokenQI) InsertToken(token *Token) error {
	return q.db.Create(token).Error
}
