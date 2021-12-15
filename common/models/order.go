package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

const (
	OrderStatusPending       = "pending"
	OrderStatusPartialFilled = "partial_filled"
	OrderStatusFullFilled    = "full_filled"
	OrderStatusCanceled      = "canceled"
)

type OrderQI interface {
	FindMarketPendingOrders(marketID string) ([]*Order, error)
	FindByAccount(trader, marketID, status string, offset, limit int) (int64, []*Order, error)
	FindByID(id string) (*Order, error)
	GetUnrealizedPNL(trader, marketID string) (decimal.Decimal, error)
	GetLockedBalance(account, tokenSymbol string, decimals int) (decimal.Decimal, error)
	InsertOrder(order *Order) error
	UpdateOrder(order *Order) error
}

type Order struct {
	ID              string          `json:"id"               db:"id" primaryKey:"true" gorm:"primary_key"`
	TraderAddress   string          `json:"trader_address"   db:"trader_address"`
	MarketID        string          `json:"market_id"        db:"market_id"`
	Side            string          `json:"side"             db:"side"`
	Price           decimal.Decimal `json:"price"            db:"price"`
	Amount          decimal.Decimal `json:"amount"           db:"amount"`
	Status          string          `json:"status"           db:"status"`
	Type            string          `json:"type"             db:"type"`
	AvailableAmount decimal.Decimal `json:"available_amount" db:"available_amount"`
	ConfirmedAmount decimal.Decimal `json:"confirmed_amount" db:"confirmed_amount"`
	CanceledAmount  decimal.Decimal `json:"canceled_amount"  db:"canceled_amount"`
	PendingAmount   decimal.Decimal `json:"pending_amount"   db:"pending_amount"`
	Signature       string          `json:"signature"        db:"signature"`
	Flags           string          `json:"flags"            db:"flags"`
	CreatedAt       time.Time       `json:"created_at"       db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"       db:"updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) AutoSetStatusByAmounts() {
	switch {
	case o.AvailableAmount.Add(o.PendingAmount).GreaterThan(decimal.Zero):
		o.Status = OrderStatusPending
	case o.ConfirmedAmount.Equal(o.Amount):
		o.Status = OrderStatusFullFilled
	case o.CanceledAmount.Equal(o.Amount):
		o.Status = OrderStatusCanceled
	default:
		o.Status = OrderStatusPartialFilled
	}
}

type orderQI struct {
	db *gorm.DB
}

func (q orderQI) FindMarketPendingOrders(marketID string) (orders []*Order, err error) {
	q.db.Where("status = 'pending' and market_id = ?", marketID).Order("created_at asc").Find(&orders)
	return
}

func (q orderQI) FindByAccount(trader, marketID, status string, offset, limit int) (int64, []*Order, error) {
	var count int64
	var orders []*Order
	statuses := strings.Split(status, ",")

	whereQuery := "trader_address = ? and market_id = ? and status IN (?)"
	err := q.db.
		Where(whereQuery, trader, marketID, statuses).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&orders).Error
	if err != nil {
		return 0, nil, err
	}

	err = q.db.Model(&Order{}).Where(whereQuery, trader, marketID, statuses).Count(&count).Error
	if err != nil {
		return 0, nil, err
	}

	return count, orders, nil
}

func (q orderQI) FindByID(id string) (*Order, error) {
	var order Order
	err := q.db.Where("id = ?", id).First(&order).Error
	return &order, err
}

func (q orderQI) GetUnrealizedPNL(trader, marketID string) (decimal.Decimal, error) {
	var unrealizedPNL decimal.NullDecimal
	query := `
		with last_price AS (
			select price
			from trades
			where market_id = 'TSLA-USDT'
			  and status = 'successful'
			order by executed_at desc
			limit 1
		)
		select sum( case when side = 'buy' THEN available_amount * ((select * from last_price) - price)
						else available_amount * (price - (select * from last_price))
						end
					) as unrealized_pnl
		from orders
		where trader_address = $1 and market_id = $2 and status = 'pending'
	`
	row := q.db.Raw(query, trader, marketID).Row()
	if err := row.Scan(&unrealizedPNL); err != nil {
		return decimal.Decimal{}, err
	}

	if !unrealizedPNL.Valid {
		return decimal.Zero, nil
	}
	return unrealizedPNL.Decimal, nil
}

func (q orderQI) GetLockedBalance(account, tokenSymbol string, decimals int) (decimal.Decimal, error) {
	var sellLockedBalance, buyLockedBalance decimal.NullDecimal

	query := `select sum(available_amount + pending_amount) as locked_balance 
			from orders where status='pending' and trader_address= $1 and market_id like $2 and side = 'sell'`
	row := q.db.Raw(query, account, tokenSymbol+"-%").Row()
	if err := row.Scan(&sellLockedBalance); err != nil {
		return decimal.Decimal{}, err
	}

	query = `select sum( (available_amount + pending_amount) * price) as locked_balance 
			from orders where status = 'pending' and trader_address = $1 and market_id like $2 and side = 'buy'`
	row = q.db.Raw(query, account, "%-"+tokenSymbol).Row()
	if err := row.Scan(&buyLockedBalance); err != nil {
		return decimal.Decimal{}, err
	}

	return sellLockedBalance.Decimal.Add(buyLockedBalance.Decimal).Mul(decimal.New(1, int32(decimals))), nil
}

func (q orderQI) InsertOrder(order *Order) error {
	return q.db.Create(order).Error
}

func (q orderQI) UpdateOrder(order *Order) error {
	return q.db.Save(order).Error
}
