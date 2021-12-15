package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

const (
	TransactionStatusCreated    = "created"
	TransactionStatusPending    = "pending"
	TransactionStatusSuccessful = "successful"
	TransactionStatusFailed     = "failed"
)

type TransactionQI interface {
	FindByID(id int64) (*Transaction, error)
	FindByHash(transactionHash string) (*Transaction, error)
	FindAllCreated() ([]*Transaction, error)
	FindAllPending() ([]*Transaction, error)
	InsertTransaction(transaction *Transaction) error
	UpdateTransaction(transaction *Transaction) error
}

type Transaction struct {
	ID          int64               `json:"id"           db:"id" primaryKey:"true"  autoIncrement:"true" gorm:"primary_key"`
	MarketID    string              `json:"market_id"    db:"market_id"`
	Status      string              `json:"status"       db:"status"`
	Hash        sql.NullString      `json:"hash"         db:"hash"`
	BlockNumber sql.NullInt64       `json:"block_number" db:"block_number"`
	GasLimit    sql.NullInt64       `json:"gas_limit"    db:"gas_limit"`
	GasUsed     sql.NullInt64       `json:"gas_used"     db:"gas_used"`
	GasPrice    decimal.NullDecimal `json:"gas_price"    db:"gas_price"`
	Nonce       sql.NullInt64       `json:"nonce"        db:"nonce"`
	Data        string              `json:"data"         db:"data"`
	ExecutedAt  time.Time           `json:"executed_at"  db:"executed_at"`
	UpdatedAt   time.Time           `json:"updated_at"   db:"updated_at"`
	CreatedAt   time.Time           `json:"created_at"   db:"created_at"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type transactionQI struct {
	db *gorm.DB
}

func (q transactionQI) FindByID(id int64) (*Transaction, error) {
	var transaction Transaction
	err := q.db.Where("id = ?", id).Find(&transaction).Error
	return &transaction, err
}

func (q transactionQI) FindByHash(transactionHash string) (*Transaction, error) {
	var transaction Transaction
	err := q.db.Where("hash = ?", transactionHash).First(&transaction).Error
	return &transaction, err
}

func (q transactionQI) FindAllCreated() ([]*Transaction, error) {
	var transactions []*Transaction
	err := q.db.Where("status = ?", TransactionStatusCreated).Order("created_at asc").Find(&transactions).Error
	return transactions, err
}

func (q transactionQI) FindAllPending() ([]*Transaction, error) {
	var transactions []*Transaction
	err := q.db.Where("status = ?", TransactionStatusPending).Order("created_at asc").Find(&transactions).Error
	return transactions, err
}

func (q transactionQI) InsertTransaction(transaction *Transaction) error {
	return q.db.Create(transaction).Error
}

func (q transactionQI) UpdateTransaction(transaction *Transaction) error {
	return q.db.Save(transaction).Error
}
