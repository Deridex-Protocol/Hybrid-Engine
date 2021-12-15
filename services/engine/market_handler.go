package engine

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"sort"
	"strings"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"bitbucket.ideasoft.io/dex/dex-backend/common/orderbook"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MarketHandler struct {
	log             *logrus.Entry
	qi              models.QI
	redis           *redis.Client
	market          *models.Market
	eventChan       chan []byte
	orderBook       *orderbook.OrderBook
	exchangeAddress ethcommon.Address
	relayerAddress  ethcommon.Address
	abiEncoder      abi.Arguments
}

func NewMarketHandler(log *logrus.Entry, qi models.QI, redis *redis.Client, market *models.Market,
	exchangeAddress, relayerAddress string) (*MarketHandler, error) {
	marketHandler := &MarketHandler{
		log:             log,
		qi:              qi,
		redis:           redis,
		market:          market,
		orderBook:       orderbook.NewOrderBook(market.ID),
		eventChan:       make(chan []byte),
		exchangeAddress: ethcommon.HexToAddress(exchangeAddress),
		relayerAddress:  ethcommon.HexToAddress(relayerAddress),
	}

	orders, err := qi.Order().FindMarketPendingOrders(market.ID)
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		if order.AvailableAmount.LessThanOrEqual(decimal.Zero) {
			continue
		}

		orderBookOrder := &orderbook.Order{
			ID:     order.ID,
			Price:  order.Price,
			Amount: order.AvailableAmount,
		}
		marketHandler.orderBook.InsertOrder(orderBookOrder, order.Side)

		if err := marketHandler.updateOrderBookSnapshot(); err != nil {
			marketHandler.log.WithError(err).Error("Failed to update order book snapshot")
			return nil, err
		}
	}

	abiEncoder, err := abi.JSON(strings.NewReader(contract.TradeMetaData.ABI))
	if err != nil {
		return nil, err
	}

	method, ok := abiEncoder.Methods["tradeExample"]
	if !ok {
		return nil, errors.New("failed to get tradeExample method from abi encoder")
	}
	marketHandler.abiEncoder = method.Inputs

	return marketHandler, nil
}

func (m *MarketHandler) Run(ctx context.Context) {
	var err error
	for {
		select {
		case data := <-m.eventChan:
			var event common.Event
			if err = json.Unmarshal(data, &event); err != nil {
				m.log.WithError(err).Error("Failed to unmarshal data to event")
				continue
			}

			switch event.Type {
			case common.EventNewOrder:
				var newOrderEvent common.NewOrderEvent
				if err = json.Unmarshal(data, &newOrderEvent); err != nil {
					m.log.WithError(err).Error("Failed to unmarshal data to new order event")
					continue
				}
				if err = m.handleNewOrder(newOrderEvent.Order); err != nil {
					m.log.WithError(err).Error("Failed to handle new order")
					continue
				}
			case common.EventCancelOrder:
				var cancelOrderEvent common.CancelOrderEvent
				if err = json.Unmarshal(data, &cancelOrderEvent); err != nil {
					m.log.WithError(err).Error("Failed to unmarshal data to cancel order event")
					continue
				}
				if err = m.handleCancelOrder(cancelOrderEvent.ID); err != nil {
					m.log.WithError(err).Error("Failed to handle cancel order")
					continue
				}
			case common.EventConfirmTransaction:
				var confirmTransactionEvent common.ConfirmTransactionEvent
				if err = json.Unmarshal(data, &confirmTransactionEvent); err != nil {
					m.log.WithError(err).Error("Failed to unmarshal data to confirm transaction event")
					continue
				}
				if err = m.handleTransactionResult(&confirmTransactionEvent); err != nil {
					m.log.WithError(err).Error("Failed to handle transaction result")
					continue
				}
			default:
				m.log.WithField("market_id", m.market.ID).
					WithField("event_type", event.Type).
					Error("Unsupported event type")
			}
		case <-ctx.Done():
			m.log.WithField("market_id", m.market.ID).Info("Market is close")
			return
		}
	}
}

func (m *MarketHandler) handleNewOrder(modelTakerOrder models.Order) error {
	m.log.WithField("market_id", modelTakerOrder.MarketID).
		WithField("type", modelTakerOrder.Type).
		WithField("side", modelTakerOrder.Side).
		WithField("price", modelTakerOrder.Price.String()).
		WithField("amount", modelTakerOrder.Amount.String()).
		Info("Handle new order")

	var wsMessages []interface{}
	var takerMatchOrder *orderbook.MatchOrder
	var makerMatchOrders []orderbook.MatchOrder
	newOrderBookOrder := orderbook.OrderFromModel(&modelTakerOrder)
	if m.orderBook.CanMatch(modelTakerOrder.Type, modelTakerOrder.Side, modelTakerOrder.Price) {
		takerMatchOrder, makerMatchOrders = m.orderBook.ExecuteOrder(newOrderBookOrder,
			modelTakerOrder.Type, modelTakerOrder.Side, m.market.AmountDecimals)

		for i := range makerMatchOrders {
			newOrderBookOrder.Amount = newOrderBookOrder.Amount.Sub(makerMatchOrders[i].MatchedAmount)

			wsMessages = append(wsMessages, common.WebSocketMessage{
				ChannelID: m.market.ID,
				Data: common.WebSocketMessageData{
					Type: common.WsUpdateOrderBookPriceType,
					Payload: common.WebsocketUpdateOrderBookPricePayload{
						Sequence: m.orderBook.Sequence,
						Side:     orderbook.ReverseSide(modelTakerOrder.Side),
						Price:    makerMatchOrders[i].Price.String(),
						Amount:   makerMatchOrders[i].MatchedAmount.String(),
					},
				},
			})

			m.log.WithField("side", orderbook.ReverseSide(modelTakerOrder.Side)).
				WithField("price", makerMatchOrders[i].Price.String()).
				WithField("amount", makerMatchOrders[i].MatchedAmount.String()).
				WithField("order_id", makerMatchOrders[i].Order.ID).
				Info("Take Liquidity")
		}
	}

	if modelTakerOrder.Type == "limit" && newOrderBookOrder.Amount.GreaterThan(decimal.Zero) {
		if m.orderBook.InsertOrder(newOrderBookOrder, modelTakerOrder.Side) {
			wsMessages = append(wsMessages, common.WebSocketMessage{
				ChannelID: m.market.ID,
				Data: common.WebSocketMessageData{
					Type: common.WsUpdateOrderBookPriceType,
					Payload: common.WebsocketUpdateOrderBookPricePayload{
						Sequence: m.orderBook.Sequence,
						Side:     modelTakerOrder.Side,
						Price:    newOrderBookOrder.Price.String(),
						Amount:   newOrderBookOrder.Amount.String(),
					},
				},
			})

			m.log.WithField("side", modelTakerOrder.Side).
				WithField("price", newOrderBookOrder.Price.String()).
				WithField("amount", newOrderBookOrder.Amount.String()).
				WithField("order_id", newOrderBookOrder.ID).
				Debug("Make Liquidity")
		}
	}

	if err := m.updateOrderBookSnapshot(); err != nil {
		return err
	}

	if takerMatchOrder != nil && makerMatchOrders != nil {
		modelMakerOrders := make(map[string]*models.Order)
		for i := range makerMatchOrders {
			modelMakerOrder, err := m.qi.Order().FindByID(makerMatchOrders[i].Order.ID)
			if err != nil {
				m.log.WithError(err).
					WithField("order_id", makerMatchOrders[i].Order.ID).
					Error("Failed to get order by id")
				return err
			}
			modelMakerOrders[makerMatchOrders[i].Order.ID] = modelMakerOrder
		}

		transactionData, err := m.createContractData(&modelTakerOrder, modelMakerOrders, makerMatchOrders)
		if err != nil {
			return err
		}

		transaction := &models.Transaction{
			Status:     models.TransactionStatusCreated,
			MarketID:   m.market.ID,
			Data:       string(transactionData),
			ExecutedAt: time.Now().UTC(),
			CreatedAt:  time.Now().UTC(),
			UpdatedAt:  time.Now().UTC(),
		}
		if err := m.qi.Transaction().InsertTransaction(transaction); err != nil {
			return err
		}

		for i := range makerMatchOrders {
			modelMakerOrder, ok := modelMakerOrders[makerMatchOrders[i].Order.ID]
			if !ok {
				continue
			}

			modelMakerOrder.AvailableAmount = modelMakerOrder.AvailableAmount.Sub(makerMatchOrders[i].MatchedAmount)
			modelTakerOrder.AvailableAmount = modelTakerOrder.AvailableAmount.Sub(makerMatchOrders[i].MatchedAmount)

			modelMakerOrder.PendingAmount = modelMakerOrder.PendingAmount.Add(makerMatchOrders[i].MatchedAmount)
			modelTakerOrder.PendingAmount = modelTakerOrder.PendingAmount.Add(makerMatchOrders[i].MatchedAmount)

			if makerMatchOrders[i].OrderIsDone {
				modelMakerOrder.CanceledAmount = modelMakerOrder.Amount.Sub(modelMakerOrder.ConfirmedAmount.Add(modelMakerOrder.PendingAmount))
				modelMakerOrder.AvailableAmount = decimal.Zero
			}

			if err := m.qi.Order().UpdateOrder(modelMakerOrder); err != nil {
				m.log.WithError(err).Error("Failed to update maker order")
				return err
			}

			trade := &models.Trade{
				TransactionID: transaction.ID,
				Status:        models.TradeStatusPending,
				MarketID:      m.market.ID,
				Maker:         modelMakerOrder.TraderAddress,
				Taker:         modelTakerOrder.TraderAddress,
				TakerSide:     modelTakerOrder.Side,
				MakerOrderID:  modelMakerOrder.ID,
				TakerOrderID:  modelTakerOrder.ID,
				Amount:        makerMatchOrders[i].MatchedAmount,
				Price:         modelMakerOrder.Price,
				CreatedAt:     time.Now().UTC(),
			}
			if err := m.qi.Trade().InsertTrade(trade); err != nil {
				return err
			}
		}

		if takerMatchOrder.OrderIsDone {
			modelTakerOrder.CanceledAmount = modelTakerOrder.Amount.Sub(modelTakerOrder.ConfirmedAmount.Add(modelTakerOrder.PendingAmount))
			modelTakerOrder.AvailableAmount = decimal.Zero
		}
	}

	if modelTakerOrder.Type == "market" {
		modelTakerOrder.CanceledAmount = modelTakerOrder.Amount.Sub(modelTakerOrder.ConfirmedAmount.Add(modelTakerOrder.PendingAmount))
		modelTakerOrder.AvailableAmount = decimal.Zero
		modelTakerOrder.AutoSetStatusByAmounts()
	}

	if err := m.qi.Order().InsertOrder(&modelTakerOrder); err != nil {
		return err
	}

	if err := m.pushMessage(wsMessages...); err != nil {
		return err
	}

	return nil
}

func (m *MarketHandler) handleCancelOrder(orderID string) error {
	m.log.WithField("order_id", orderID).Info("Handle cancel order")

	order, err := m.qi.Order().FindByID(orderID)
	if err != nil {
		return err
	}

	if m.orderBook.RemoveOrder(order.ID, order.Side, order.Price) {
		if err := m.updateOrderBookSnapshot(); err != nil {
			return err
		}

		msg := common.WebSocketMessage{
			ChannelID: m.market.ID,
			Data: common.WebSocketMessageData{
				Type: common.WsUpdateOrderBookPriceType,
				Payload: common.WebsocketUpdateOrderBookPricePayload{
					Sequence: m.orderBook.Sequence,
					Side:     order.Side,
					Price:    order.Price.String(),
					Amount:   order.Amount.String(),
				},
			},
		}
		if err := m.pushMessage(msg); err != nil {
			return err
		}
	}

	order.CanceledAmount = order.CanceledAmount.Add(order.AvailableAmount)
	order.AvailableAmount = decimal.Zero
	order.AutoSetStatusByAmounts()

	return m.qi.Order().UpdateOrder(order)
}

func (m *MarketHandler) handleTransactionResult(event *common.ConfirmTransactionEvent) error {
	m.log.WithField("tx_hash", event.Hash).Info("Handle transaction result")

	transaction, err := m.qi.Transaction().FindByHash(event.Hash)
	if err != nil {
		return err
	}

	transaction.Status = event.Status
	transaction.ExecutedAt = time.Unix(int64(event.Timestamp), 0)

	if err := m.qi.Transaction().UpdateTransaction(transaction); err != nil {
		return err
	}

	trades, err := m.qi.Trade().FindTradesByHash(event.Hash)
	if err != nil {
		return err
	}

	if len(trades) == 0 {
		return nil
	}

	takerOrder, err := m.qi.Order().FindByID(trades[0].TakerOrderID)
	if err != nil {
		return err
	}

	var wsMessages []interface{}

	for _, trade := range trades {
		makerOrder, err := m.qi.Order().FindByID(trade.MakerOrderID)
		if err != nil {
			return err
		}

		takerOrder.PendingAmount = takerOrder.PendingAmount.Sub(trade.Amount)
		makerOrder.PendingAmount = makerOrder.PendingAmount.Sub(trade.Amount)

		switch event.Status {
		case models.TradeStatusFailed:
			takerOrder.CanceledAmount = takerOrder.CanceledAmount.Add(trade.Amount)
			makerOrder.CanceledAmount = makerOrder.CanceledAmount.Add(trade.Amount)
		case models.TradeStatusSuccessful:
			takerOrder.ConfirmedAmount = takerOrder.ConfirmedAmount.Add(trade.Amount)
			makerOrder.ConfirmedAmount = makerOrder.ConfirmedAmount.Add(trade.Amount)
		}

		makerOrder.AutoSetStatusByAmounts()

		if err = m.qi.Order().UpdateOrder(makerOrder); err != nil {
			return err
		}

		trade.Status = event.Status
		trade.ExecutedAt = time.Unix(int64(event.Timestamp), 0)

		if err := m.qi.Trade().UpdateTrade(trade); err != nil {
			return err
		}

		if trade.Status == models.TradeStatusSuccessful {
			wsMessages = append(wsMessages, &common.WebSocketMessage{
				ChannelID: m.market.ID,
				Data: common.WebSocketMessageData{
					Type: common.WsMarketNewTradeType,
					Payload: common.WebsocketMarketNewTradePayload{
						Amount:     trade.Amount.String(),
						Price:      trade.Price.String(),
						ExecutedAt: trade.ExecutedAt,
					},
				},
			})
		}
	}

	if err := m.pushMessage(wsMessages...); err != nil {
		return err
	}

	takerOrder.AutoSetStatusByAmounts()

	return m.qi.Order().UpdateOrder(takerOrder)
}

func (m *MarketHandler) createContractData(modelTakerOrder *models.Order, modelMakerOrders map[string]*models.Order,
	makerOrders []orderbook.MatchOrder) ([]byte, error) {
	var accounts = []ethcommon.Address{m.relayerAddress}
	var accountsMap = map[string]struct{}{m.relayerAddress.Hex(): {}}

	accounts = append(accounts, ethcommon.HexToAddress(modelTakerOrder.TraderAddress))
	accountsMap[modelTakerOrder.TraderAddress] = struct{}{}

	tradesDataMap := make(map[string][][]byte)

	for i := range makerOrders {
		modelMakerOrder, ok := modelMakerOrders[makerOrders[i].Order.ID]
		if !ok {
			continue
		}

		if _, ok := accountsMap[modelMakerOrder.TraderAddress]; !ok {
			accountsMap[modelMakerOrder.TraderAddress] = struct{}{}
			accounts = append(accounts, ethcommon.HexToAddress(modelMakerOrder.TraderAddress))
		}

		fillAmount := makerOrders[i].MatchedAmount.Mul(decimal.New(1, int32(m.market.AmountDecimals)))
		fillPrice := modelMakerOrder.Price.Mul(decimal.New(1, int32(m.market.PriceDecimals)))

		makerData, err := m.createContractTradeData(modelMakerOrder, fillAmount, fillPrice)
		if err != nil {
			return nil, err
		}
		tradesDataMap[modelMakerOrder.TraderAddress] = append(tradesDataMap[modelMakerOrder.TraderAddress], makerData)

		takerData, err := m.createContractTradeData(modelTakerOrder, fillAmount, fillPrice)
		if err != nil {
			return nil, err
		}
		tradesDataMap[modelTakerOrder.TraderAddress] = append(tradesDataMap[modelTakerOrder.TraderAddress], takerData)
	}

	// sort accounts slice for contract
	sort.SliceStable(accounts, func(i, j int) bool {
		return accounts[i].String() < accounts[j].String()
	})

	var takerAddressIndex *big.Int
	for i := range accounts {
		if accounts[i].Hex() == m.relayerAddress.Hex() {
			takerAddressIndex = big.NewInt(int64(i))
			break
		}
	}

	var trades []contract.DxlnTradeTradeArg
	for makerAddress, tradesData := range tradesDataMap {
		var makerAddressIndex *big.Int
		for i := range accounts {
			if accounts[i].Hex() == ethcommon.HexToAddress(makerAddress).Hex() {
				makerAddressIndex = big.NewInt(int64(i))
				break
			}
		}

		for i := range tradesData {
			trades = append(trades, contract.DxlnTradeTradeArg{
				TakerIndex: takerAddressIndex,
				MakerIndex: makerAddressIndex,
				Trader:     m.exchangeAddress,
				Data:       tradesData[i],
			})
		}
	}

	transactionData, err := json.Marshal(common.TransactionData{
		Accounts: accounts,
		Trades:   trades,
	})
	if err != nil {
		return nil, err
	}

	return transactionData, nil
}

func (m *MarketHandler) createContractTradeData(modelOrder *models.Order, fillAmount, fillPrice decimal.Decimal) ([]byte, error) {
	amount := modelOrder.Amount.Mul(decimal.New(1, int32(m.market.AmountDecimals)))
	tradeData := contract.DxlnOrdersTradeData{
		Order: contract.DxlnOrdersOrder{
			Amount:       amount.BigInt(),
			LimitPrice:   modelOrder.Price.Mul(decimal.New(1, int32(m.market.PriceDecimals))).BigInt(),
			TriggerPrice: big.NewInt(0),
			LimitFee:     amount.Mul(m.market.TakerFeeRate).BigInt(),
			Maker:        ethcommon.HexToAddress(modelOrder.TraderAddress),
			Taker:        m.relayerAddress,
			Expiration:   big.NewInt(0),
		},
		Fill: contract.DxlnOrdersFill{
			Amount:        fillAmount.BigInt(),
			Price:         fillPrice.BigInt(),
			Fee:           fillAmount.Mul(m.market.TakerFeeRate).BigInt(),
			IsNegativeFee: false,
		},
	}

	flags, err := hex.DecodeString(modelOrder.Flags)
	if err != nil {
		return nil, err
	}
	copy(tradeData.Order.Flags[:], flags[:32])

	sign, err := hex.DecodeString(modelOrder.Signature)
	if err != nil {
		return nil, err
	}
	sign = append(sign, common.SignatureTypeDecimal)

	copy(tradeData.Signature.R[:], sign[:32])
	copy(tradeData.Signature.S[:], sign[32:64])
	copy(tradeData.Signature.VType[:], sign[64:])

	return m.abiEncoder.Pack(tradeData)
}

func (m *MarketHandler) pushMessage(messages ...interface{}) error {
	for i := range messages {
		msgBytes, err := json.Marshal(messages[i])
		if err != nil {
			m.log.WithError(err).Error("Failed to marshal websocket message")
			return err
		}

		m.log.WithField("msg_json", string(msgBytes)).Info("Send message to websocket")

		if err := m.redis.LPush(common.WebsocketMessageQueueKey, msgBytes).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketHandler) updateOrderBookSnapshot() error {
	snapshotBytes, err := json.Marshal(m.orderBook.Snapshot())
	if err != nil {
		return err
	}
	if err = m.redis.Set(common.OrderBookSnapshotKey+m.market.ID, string(snapshotBytes), 0).Err(); err != nil {
		return err
	}
	return nil
}
