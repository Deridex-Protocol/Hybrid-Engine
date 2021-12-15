package orderbook

import (
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID     string
	Price  decimal.Decimal
	Amount decimal.Decimal
}

type MatchOrder struct {
	Order
	OrderIsDone   bool
	MatchedAmount decimal.Decimal
}

type Snapshot struct {
	Sequence uint64      `json:"sequence"`
	Bids     [][2]string `json:"bids"`
	Asks     [][2]string `json:"asks"`
}

func OrderFromModel(o *models.Order) *Order {
	return &Order{
		ID:     o.ID,
		Price:  o.Price,
		Amount: o.Amount,
	}
}

func ReverseSide(side string) string {
	if side == "buy" {
		return "sell"
	}
	return "buy"
}
