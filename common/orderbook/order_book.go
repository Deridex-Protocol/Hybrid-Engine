package orderbook

import (
	"strings"
	"sync"

	"github.com/petar/GoLLRB/llrb"
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	Sequence uint64
	mutex    sync.RWMutex
	marketID string
	bidsTree *llrb.LLRB
	asksTree *llrb.LLRB
}

// NewOrderBook return a new order book
// asks - sell orders, bids - buy orders
func NewOrderBook(marketID string) *OrderBook {
	return &OrderBook{
		marketID: marketID,
		bidsTree: llrb.New(),
		asksTree: llrb.New(),
	}
}

func (o *OrderBook) Snapshot() *Snapshot {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	asks := make([][2]string, 0, 0)
	bids := make([][2]string, 0, 0)

	var asyncWaitGroup sync.WaitGroup
	asyncWaitGroup.Add(1)
	go func() {
		o.asksTree.AscendGreaterOrEqual(newPriceLevel(decimal.Zero), func(item llrb.Item) bool {
			if priceLevel, ok := item.(*PriceLevel); ok {
				asks = append(asks, [2]string{priceLevel.price.String(), priceLevel.totalAmount.String()})
			}
			return true
		})
		asyncWaitGroup.Done()
	}()

	asyncWaitGroup.Add(1)
	go func() {
		o.bidsTree.DescendLessOrEqual(newPriceLevel(decimal.New(1, 99)), func(i llrb.Item) bool {
			if priceLevel, ok := i.(*PriceLevel); ok {
				bids = append(bids, [2]string{priceLevel.price.String(), priceLevel.totalAmount.String()})
			}
			return true
		})
		asyncWaitGroup.Done()
	}()

	asyncWaitGroup.Wait()

	return &Snapshot{
		Bids:     bids,
		Asks:     asks,
		Sequence: o.Sequence,
	}
}

func (o *OrderBook) InsertOrder(order *Order, side string) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	tree := o.getTreeBySide(side)

	priceLevelItem := tree.Get(newPriceLevel(order.Price))
	var priceLevel *PriceLevel
	if priceLevelItem != nil {
		var ok bool
		priceLevel, ok = priceLevelItem.(*PriceLevel)
		if !ok {
			return false
		}
		priceLevel.InsertOrder(order)
	} else {
		price := newPriceLevel(order.Price)
		price.InsertOrder(order)
		tree.InsertNoReplace(price)
	}

	o.Sequence = o.Sequence + 1

	return true
}

func (o *OrderBook) RemoveOrder(orderID, side string, price decimal.Decimal) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	tree := o.getTreeBySide(side)

	priceLevelItem := tree.Get(newPriceLevel(price))
	if priceLevelItem == nil {
		return false
	}
	priceLevel, ok := priceLevelItem.(*PriceLevel)
	if !ok {
		return false
	}

	priceLevel.RemoveOrder(orderID)
	if len(priceLevel.orderMap) == 0 {
		tree.Delete(priceLevel)
	}

	o.Sequence = o.Sequence + 1

	return true
}

func (o *OrderBook) ChangeOrder(order *Order, side string) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	tree := o.getTreeBySide(side)

	priceLevelItem := tree.Get(newPriceLevel(order.Price))
	if priceLevelItem == nil {
		return false
	}
	priceLevel, ok := priceLevelItem.(*PriceLevel)
	if !ok {
		return false
	}

	priceLevel.UpdateOrderAmount(order.ID, order.Amount)

	o.Sequence = o.Sequence + 1

	return true
}

func (o *OrderBook) CanMatch(orderType, orderSide string, price decimal.Decimal) bool {
	o.mutex.RLock()
	defer o.mutex.RUnlock()

	if strings.EqualFold(orderType, "market") {
		if orderSide == "buy" {
			return o.asksTree.Len() > 0
		} else {
			return o.bidsTree.Len() > 0
		}
	}

	if strings.EqualFold("buy", orderSide) {
		minItem := o.asksTree.Min()
		if minItem == nil {
			return false
		}
		priceLevel, ok := minItem.(*PriceLevel)
		if !ok {
			return false
		}

		return price.GreaterThanOrEqual(priceLevel.price)
	} else {
		maxItem := o.bidsTree.Max()
		if maxItem == nil {
			return false
		}
		priceLevel, ok := maxItem.(*PriceLevel)
		if !ok {
			return false
		}

		return price.LessThanOrEqual(priceLevel.price)
	}
}

// ExecuteOrder modify order book
func (o *OrderBook) ExecuteOrder(order *Order, orderType, orderSide string, marketAmountDecimals int) (*MatchOrder, []MatchOrder) {
	if order.Amount.LessThanOrEqual(decimal.Zero) {
		return nil, nil
	}

	var makerOrders []Order
	var makerMatchOrders []MatchOrder
	var leftAmount = order.Amount

	// This function will be called multi times
	// Return false to break the loop
	limitOrderIterator := func(item llrb.Item) bool {
		itemPrice, ok := item.(*PriceLevel)
		if !ok {
			return false
		}

		// if buy and order book price > taker price return false
		// if sell and order book price < taker price return false
		if orderSide == "buy" && itemPrice.price.GreaterThan(order.Price) {
			return false
		} else if orderSide == "sell" && itemPrice.price.LessThan(order.Price) {
			return false
		}

		for _, priceOrder := range itemPrice.orderMap {
			makerMatchOrder := MatchOrder{Order: *priceOrder}
			if leftAmount.GreaterThanOrEqual(priceOrder.Amount) {
				makerMatchOrder.OrderIsDone = true
				makerMatchOrder.MatchedAmount = priceOrder.Amount
				leftAmount = leftAmount.Sub(priceOrder.Amount)
			} else {
				makerMatchOrder.MatchedAmount = leftAmount
				leftAmount = decimal.Zero
			}

			makerOrders = append(makerOrders, *priceOrder)
			makerMatchOrders = append(makerMatchOrders, makerMatchOrder)

			if leftAmount.LessThanOrEqual(decimal.Zero) {
				return false
			}
		}

		return leftAmount.GreaterThan(decimal.Zero)
	}

	marketOrderIterator := func(item llrb.Item) bool {
		itemPrice, ok := item.(*PriceLevel)
		if !ok {
			return false
		}

		for _, priceOrder := range itemPrice.orderMap {
			makerMatchOrder := MatchOrder{Order: *priceOrder}

			if leftAmount.GreaterThanOrEqual(priceOrder.Amount) {
				makerMatchOrder.MatchedAmount = priceOrder.Amount
				makerMatchOrder.OrderIsDone = true
				leftAmount = leftAmount.Sub(priceOrder.Amount)
			} else {
				makerMatchOrder.MatchedAmount = leftAmount
				leftAmount = decimal.Zero
			}

			makerOrders = append(makerOrders, *priceOrder)
			makerMatchOrders = append(makerMatchOrders, makerMatchOrder)

			if leftAmount.LessThanOrEqual(decimal.Zero) {
				return false
			}
		}

		return leftAmount.GreaterThan(decimal.Zero)
	}

	iterator := limitOrderIterator
	if orderType == "market" {
		iterator = marketOrderIterator
	}

	o.mutex.Lock()
	if orderSide == "sell" {
		o.bidsTree.DescendLessOrEqual(newPriceLevel(decimal.New(1, 99)), iterator)
	} else {
		o.asksTree.AscendGreaterOrEqual(newPriceLevel(decimal.Zero), iterator)
	}
	o.mutex.Unlock()

	makerSide := ReverseSide(orderSide)

	for i := range makerOrders {
		if makerMatchOrders[i].OrderIsDone {
			o.RemoveOrder(makerOrders[i].ID, makerSide, makerOrders[i].Price)
		} else {
			makerOrders[i].Amount = makerOrders[i].Amount.Sub(makerMatchOrders[i].MatchedAmount)
			o.ChangeOrder(&makerOrders[i], makerSide)
		}
	}

	takerMatchOrder := &MatchOrder{
		Order: Order{
			ID:     order.ID,
			Price:  order.Price,
			Amount: order.Amount,
		},
		OrderIsDone:   leftAmount.LessThanOrEqual(decimal.Zero),
		MatchedAmount: order.Amount.Sub(leftAmount),
	}
	return takerMatchOrder, makerMatchOrders
}

func (o *OrderBook) getTreeBySide(side string) *llrb.LLRB {
	if side == "sell" {
		return o.asksTree
	}
	return o.bidsTree
}
