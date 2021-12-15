package handlers

import (
	"encoding/json"
	"net/http"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type GetOrdersHandler struct {
	qi    models.QI
	log   *logrus.Entry
	redis *redis.Client
}

func NewGetOrderHandler(qi models.QI, log *logrus.Entry, redis *redis.Client) GetOrdersHandler {
	return GetOrdersHandler{
		qi:    qi,
		log:   log,
		redis: redis,
	}
}

// GetSingleOrder returns single order by order id
// @Summary Get single order by order id
// @Description Returns single order by order id. Response contains universal data structure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Orders
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} api.Response{data=models.Order} "Order data is inside data field"
// @Router /orders/{order_id} [get]
func (h *GetOrdersHandler) GetSingleOrder(p interface{}) (interface{}, error) {
	req := p.(*QuerySingleOrderReq)

	order, err := h.qi.Order().FindByID(req.OrderID)
	if err != nil && err != gorm.ErrRecordNotFound {
		h.log.WithError(err).Error("Failed to get order by id")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	return order, nil
}

// GetOrders returns account orders of the market
// @Summary Get all account orders
// @Description Returns all account orders. Response contains universal data structure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Param market_id query string true "Market ID"
// @Param status query string false "Status or statuses separated by comma"
// @Param page query int false "Page"
// @Param perPage query int false "Rows per page"
// @Tags Orders
// @Produce json
// @Success 200 {object} api.Response{data=handlers.OrdersResp} "Account orders are inside data field"
// @Router /orders [get]
func (h *GetOrdersHandler) GetOrders(p interface{}) (interface{}, error) {
	req := p.(*QueryOrderReq)
	if req.Status == "" {
		req.Status = models.OrderStatusPending
	}
	if req.PerPage <= 0 {
		req.PerPage = 20
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	offset := req.PerPage * (req.Page - 1)
	limit := req.PerPage

	count, orders, err := h.qi.Order().FindByAccount(req.Address, req.MarketID, req.Status, offset, limit)
	if err != nil {
		h.log.WithError(err).Error("Failed to get orders by account")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return OrdersResp{
		Count:  count,
		Orders: orders,
	}, nil
}

// GetOrdersInfo returns account orders info of the market
// @Summary Get all account orders info
// @Description Returns all account orders info. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Param market_id query string true "Market ID"
// @Tags Orders
// @Produce json
// @Success 200 {object} api.Response{data=handlers.GetOrdersInfoResp} "Account orders info are inside data field"
// @Router /orders/info [get]
func (h *GetOrdersHandler) GetOrdersInfo(p interface{}) (interface{}, error) {
	req := p.(*GetOrdersInfoReq)

	unrealizedPNL, err := h.qi.Order().GetUnrealizedPNL(req.Address, req.MarketID)
	if err != nil {
		return nil, err
	}

	return &GetOrdersInfoResp{
		UnrealizedPNL: unrealizedPNL,
	}, nil
}

// CancelOrder cancels order placed on the market
// @Summary Cancel placed order
// @Description Cancels order placed on the market. Response contains universal datastructure. 'data' field contains response value, in case of errors - 'status' and 'desc' contain description of the error, and the 'data' value is empty. This request requires authentication
// @Tags Orders
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} api.Response "Contains result of order cancellation"
// @Router /orders/{order_id} [delete]
func (h *GetOrdersHandler) CancelOrder(p interface{}) (interface{}, error) {
	req := p.(*CancelOrderReq)
	order, err := h.qi.Order().FindByID(req.ID)
	if err != nil {
		h.log.WithError(err).Error("Failed to get order by id")
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	if order.Status != models.OrderStatusPending {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Order is not pending")
	}

	cancelOrderEvent := common.CancelOrderEvent{
		Type:     common.EventCancelOrder,
		MarketID: order.MarketID,
		ID:       order.ID,
	}

	cancelOrderEventJSON, err := json.Marshal(cancelOrderEvent)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError)
	}

	return nil, h.redis.LPush(common.EngineEventQueueKey, cancelOrderEventJSON).Err()
}
