package api

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/contract"
	"bitbucket.ideasoft.io/dex/dex-backend/common/models"
	_ "bitbucket.ideasoft.io/dex/dex-backend/services/api/docs"
	"bitbucket.ideasoft.io/dex/dex-backend/services/api/handlers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Service struct {
	log            *logrus.Entry
	cfg            *Config
	qi             models.QI
	redis          *redis.Client
	ethClient      *ethclient.Client
	contractClient *contract.Contract
}

func NewApiService(log *logrus.Entry, cfg *Config, qi models.QI, redis *redis.Client, ethClient *ethclient.Client,
	contractClient *contract.Contract) Service {
	return Service{
		log:            log,
		cfg:            cfg,
		qi:             qi,
		redis:          redis,
		ethClient:      ethClient,
		contractClient: contractClient,
	}
}

func (s *Service) Run(ctx context.Context) {
	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = echoErrorHandler
	e.Use(middleware.Logger())
	e.Use(s.recoverHandler)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, common.AuthenticationHeaderKey},
	}))

	s.loadRoutes(e)

	server := &http.Server{
		Addr:         ":3001",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		if err := e.StartServer(server); err != nil {
			e.Logger.Info("shutting down the server: %v", err)
			return
		}
	}()

	<-ctx.Done()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// @title Deridex Backend API
// @version 1.0.0
// @description This is a documentation for the Deridex market backend API.

// @contact.name API Support
// @contact.url https://deridex-dev.ml/support
// @contact.email support@deridex-dev.ml

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host deridex-dev.ml
// @BasePath /api/v1
func (s *Service) loadRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")

	g.GET("/swagger/*", echoSwagger.WrapHandler)

	// markets
	marketHandler := handlers.NewMarketHandler(s.qi, s.log, s.redis)
	g.Add("GET", "/markets", commonHandler(nil, marketHandler.GetMarkets))
	g.Add("GET", "/markets/:market_id/orderbook", commonHandler(&handlers.OrderBookReq{}, marketHandler.GetOrderBook))

	// trades
	tradeHandler := handlers.NewTradeHandler(s.qi, s.log)
	g.Add("GET", "/markets/:market_id/trades", commonHandler(&handlers.QueryTradeReq{}, tradeHandler.GetAllTrades))
	g.Add("GET", "/markets/:market_id/trades/mine", commonHandler(&handlers.QueryTradeReq{}, tradeHandler.GetAccountTrades), authMiddleware)
	g.Add("GET", "/markets/:market_id/candles", commonHandler(&handlers.CandlesReq{}, tradeHandler.GetTradingView))

	// get order
	getOrderHandler := handlers.NewGetOrderHandler(s.qi, s.log, s.redis)
	g.Add("GET", "/orders", commonHandler(&handlers.QueryOrderReq{}, getOrderHandler.GetOrders), authMiddleware)
	g.Add("GET", "/orders/:order_id", commonHandler(&handlers.QuerySingleOrderReq{}, getOrderHandler.GetSingleOrder), authMiddleware)
	g.Add("GET", "/orders/info", commonHandler(&handlers.GetOrdersInfoReq{}, getOrderHandler.GetOrdersInfo), authMiddleware)
	g.Add("DELETE", "/orders/:order_id", commonHandler(&handlers.CancelOrderReq{}, getOrderHandler.CancelOrder), authMiddleware)

	// create order
	createOrderHandler := handlers.NewCreateOrderHandler(s.qi, s.log, s.redis, s.contractClient, s.cfg.ChanID, s.cfg.ExchangeAddress, s.cfg.RelayerAddress)
	g.Add("POST", "/orders/build/limit", commonHandler(&handlers.BuildLimitOrderReq{}, createOrderHandler.BuildLimitOrder), authMiddleware)
	g.Add("POST", "/orders/build/market", commonHandler(&handlers.BuildMarketOrderReq{}, createOrderHandler.BuildMarketOrder), authMiddleware)
	g.Add("POST", "/orders", commonHandler(&handlers.PlaceOrderReq{}, createOrderHandler.PlaceOrder), authMiddleware)
}

func echoErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var desc string
	var code int

	e := c.Echo()

	if httpError, ok := err.(*echo.HTTPError); ok {
		desc = httpError.Error()
		code = httpError.Code
	} else if errors, ok := err.(validator.ValidationErrors); ok {
		var buff bytes.Buffer
		for _, err := range errors {
			switch err.Tag() {
			case "required":
				buff.WriteString(fmt.Sprintf("%s is required", err.Field()))
			default:
				buff.WriteString(fmt.Sprintf("Key: '%s' Error:Field validation for '%s' failed on the '%s' tag", err.Namespace(), err.Field(), err.Tag()))
			}
			buff.WriteString(";")
		}
		desc = buff.String()
	} else if e.Debug {
		desc = err.Error()
	} else {
		desc = "something wrong"
	}

	// Send response
	if err = c.JSON(code, Response{Desc: desc}); err != nil {
		e.Logger.Error(err)
	}
}

func (s Service) recoverHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 2048)
				length := runtime.Stack(stack, false)

				s.log.WithError(err).
					WithField("stack", stack[:length]).
					Error("Unhandled error")
				c.Error(err)
			}
		}()
		return next(c)
	}
}
