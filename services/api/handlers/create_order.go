package handlers

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	commoncrypto "bitbucket.ideasoft.io/dex/dex-backend/common/crypto"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const OrderCacheKey = "OrderCache:"

type CreateOrdersHandler struct {
	qi              models.QI
	log             *logrus.Entry
	redis           *redis.Client
	contractClient  *contract.Contract
	chanID          int64
	exchangeAddress ethcommon.Address
	relayerAddress  ethcommon.Address
}

func NewCreateOrderHandler(qi models.QI, log *logrus.Entry, redis *redis.Client, cc *contract.Contract,
	chanID int64, exchangeAddress, relayerAddress string) CreateOrdersHandler {
	return CreateOrdersHandler{
		qi:              qi,
		log:             log,
		redis:           redis,
		contractClient:  cc,
		chanID:          chanID,
		exchangeAddress: ethcommon.HexToAddress(exchangeAddress),
		relayerAddress:  ethcommon.HexToAddress(relayerAddress),
	}
}

// BuildLimitOrder builds and returns new limit order
// @Summary Build a new limit order
// @Description Builds and returns new limit order. Response contains universal data structure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Orders
// @Produce json
// @Param order_params body handlers.BuildLimitOrderReq true "Contains parameters of new order"
// @Success 200 {object} api.Response{data=handlers.BuildOrderResp} "Order data is inside data field"
// @Router /orders/build/limit [post]
func (h *CreateOrdersHandler) BuildLimitOrder(p interface{}) (interface{}, error) {
	req := p.(*BuildLimitOrderReq)

	market, err := h.qi.Market().FindMarketByID(req.MarketID)
	if err != nil {
		h.log.WithError(err).
			WithField("market_id", req.MarketID).
			Error("Failed to get market by id")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	price, err := decimal.NewFromString(req.Price)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if price.LessThanOrEqual(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_price")
	}
	if !price.Mod(decimal.New(1, int32(-1*market.PriceDecimals))).Equal(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_price_unit")
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_amount")
	}
	if !amount.Mod(decimal.New(1, int32(-1*market.AmountDecimals))).Equal(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_amount_unit")
	}

	if amount.Mul(price).LessThan(market.MinOrderSize) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "order_less_than_minOrderSize")
	}

	buildOrder := &BuildOrder{
		MarketID:  req.MarketID,
		Side:      req.Side,
		Price:     req.Price,
		Amount:    req.Amount,
		OrderType: "limit",
	}

	orderID, err := h.BuildAndCacheOrder(req.Address, buildOrder)
	if err != nil {
		h.log.WithError(err).Error("Failed get build and cache order")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return BuildOrderResp{OrderID: orderID}, nil
}

// BuildMarketOrder builds and returns new market order
// @Summary Build a new market order
// @Description Builds and returns new market order. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Orders
// @Produce json
// @Param order_params body handlers.BuildMarketOrderReq true "Contains parameters of new order"
// @Success 200 {object} api.Response{data=handlers.BuildOrderResp} "Order data is inside data field"
// @Router /orders/build/market [post]
func (h *CreateOrdersHandler) BuildMarketOrder(p interface{}) (interface{}, error) {
	req := p.(*BuildMarketOrderReq)

	market, err := h.qi.Market().FindMarketByID(req.MarketID)
	if err != nil {
		h.log.WithError(err).
			WithField("market_id", req.MarketID).
			Error("Failed to get market by id")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_amount")
	}
	if !amount.Mod(decimal.New(1, int32(-1*market.AmountDecimals))).Equal(decimal.Zero) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid_amount_unit")
	}

	buildOrder := &BuildOrder{
		MarketID:  req.MarketID,
		Side:      req.Side,
		Price:     decimal.Zero.String(),
		Amount:    req.Amount,
		OrderType: "market",
	}

	if req.Side == "buy" {
		buildOrder.Price = decimal.New(1, 13).String()
	}

	orderID, err := h.BuildAndCacheOrder(req.Address, buildOrder)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return BuildOrderResp{OrderID: orderID}, nil
}

func (h *CreateOrdersHandler) BuildAndCacheOrder(address string, order *BuildOrder) (string, error) {
	market, err := h.qi.Market().FindMarketByID(order.MarketID)
	if err != nil {
		h.log.WithError(err).
			WithField("market_id", order.MarketID).
			Error("Failed to get market by id")
		return "", echo.NewHTTPError(http.StatusInternalServerError)
	}

	price, err := decimal.NewFromString(order.Price)
	if err != nil {
		return "", err
	}

	amount, err := decimal.NewFromString(order.Amount)
	if err != nil {
		return "", err
	}

	flags, err := common.GetTradeFlags(order.Side == "buy")
	if err != nil {
		return "", err
	}

	dxlnOrderAmount := amount.Mul(decimal.New(1, int32(market.AmountDecimals)))
	dxlnOrder := &contract.DxlnOrdersOrder{
		Flags:        flags,
		Amount:       dxlnOrderAmount.BigInt(),
		LimitPrice:   price.Mul(decimal.New(1, int32(market.PriceDecimals))).BigInt(),
		TriggerPrice: big.NewInt(0),
		LimitFee:     dxlnOrderAmount.Mul(market.TakerFeeRate).BigInt(),
		Maker:        ethcommon.HexToAddress(address),
		Taker:        h.relayerAddress,
		Expiration:   big.NewInt(0),
	}

	orderHash, err := commoncrypto.GetOrderHash(dxlnOrder, h.chanID, h.exchangeAddress)
	if err != nil {
		return "", err
	}

	cacheOrder := CacheOrder{
		ID:       "0x" + hex.EncodeToString(orderHash),
		MarketID: order.MarketID,
		Side:     order.Side,
		Type:     order.OrderType,
		Price:    price,
		Amount:   amount,
		Address:  address,
		Flags:    flags,
	}

	orderJSON, err := json.Marshal(cacheOrder)
	if err != nil {
		return "", err
	}

	// Cache the build order for 60 seconds, if we still not get signature in the period. The order will be dropped.
	if err := h.redis.Set(OrderCacheKey+cacheOrder.ID, string(orderJSON), time.Second*60).Err(); err != nil {
		return "", err
	}

	return cacheOrder.ID, err
}

// PlaceOrder places existing order on the market
// @Summary Place existing order
// @Description Places existing order on the market. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Orders
// @Produce json
// @Param order_params body handlers.PlaceOrderReq true "Contains parameters of existing order"
// @Success 200 {object} api.Response "Contains result of order placing"
// @Router /orders [post]
func (h *CreateOrdersHandler) PlaceOrder(p interface{}) (interface{}, error) {
	req := p.(*PlaceOrderReq)

	if strings.HasPrefix(req.OrderID, "0x") || strings.HasPrefix(req.OrderID, "0X") {
		req.OrderID = req.OrderID[2:]
	}
	orderHash, err := hex.DecodeString(req.OrderID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if strings.HasPrefix(req.Signature, "0x") || strings.HasPrefix(req.Signature, "0X") {
		req.Signature = req.Signature[2:]
	}
	signatureBytes, err := hex.DecodeString(req.Signature)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	valid, err := commoncrypto.IsValidOrderSignature(req.Address, orderHash, signatureBytes)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if !valid {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "bad signature")
	}

	cacheOrderStr, err := h.redis.Get(OrderCacheKey + "0x" + req.OrderID).Result()
	if err != nil {
		h.log.WithError(err).Error("Failed to get cache order")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	var cacheOrder CacheOrder
	if err = json.Unmarshal([]byte(cacheOrderStr), &cacheOrder); err != nil {
		h.log.WithError(err).Error("Failed to unmarshal order")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	order := models.Order{
		ID:              "0x" + req.OrderID,
		TraderAddress:   req.Address,
		MarketID:        cacheOrder.MarketID,
		Side:            cacheOrder.Side,
		Price:           cacheOrder.Price,
		Amount:          cacheOrder.Amount,
		Status:          models.OrderStatusPending,
		Type:            cacheOrder.Type,
		AvailableAmount: cacheOrder.Amount,
		ConfirmedAmount: decimal.Zero,
		CanceledAmount:  decimal.Zero,
		PendingAmount:   decimal.Zero,
		Signature:       req.Signature,
		Flags:           hex.EncodeToString(cacheOrder.Flags[:]),
		CreatedAt:       time.Now().UTC(),
	}

	newOrderEvent, err := json.Marshal(common.NewOrderEvent{
		Type:     common.EventNewOrder,
		MarketID: cacheOrder.MarketID,
		Order:    order,
	})
	if err != nil {
		h.log.WithError(err).Error("Failed to marshal order event")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	if err := h.redis.LPush(common.EngineEventQueueKey, newOrderEvent).Err(); err != nil {
		h.log.WithError(err).Error("Failed to push to redis engine event")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return nil, nil
}
