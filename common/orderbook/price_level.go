package orderbook

import (
	"github.com/petar/GoLLRB/llrb"
	"github.com/shopspring/decimal"
)

type PriceLevel struct {
	price       decimal.Decimal
	totalAmount decimal.Decimal
	orderMap    map[string]*Order
}

func newPriceLevel(price decimal.Decimal) *PriceLevel {
	return &PriceLevel{
		price:       price,
		totalAmount: decimal.Zero,
		orderMap:    make(map[string]*Order),
	}
}

func (p *PriceLevel) Less(item llrb.Item) bool {
	another := item.(*PriceLevel)
	return p.price.LessThan(another.price)
}

func (p *PriceLevel) InsertOrder(newOrder *Order) {
	if _, ok := p.orderMap[newOrder.ID]; ok {
		return
	}
	p.orderMap[newOrder.ID] = newOrder
	p.totalAmount = p.totalAmount.Add(newOrder.Amount)
}

func (p *PriceLevel) UpdateOrderAmount(id string, newAmount decimal.Decimal) {
	if oldOrder, ok := p.orderMap[id]; ok {
		p.totalAmount = p.totalAmount.Sub(oldOrder.Amount).Add(newAmount)
		p.orderMap[id].Amount = newAmount
	}
}

func (p *PriceLevel) RemoveOrder(id string) {
	if order, ok := p.orderMap[id]; ok {
		delete(p.orderMap, order.ID)
		p.totalAmount = p.totalAmount.Sub(order.Amount)
	}
}
