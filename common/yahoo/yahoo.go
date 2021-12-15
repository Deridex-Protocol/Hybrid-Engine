package yahoo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

const yahooFinancePriceURL = "https://query1.finance.yahoo.com/v10/finance/quoteSummary"

const (
	MarketStatePre     = "PRE"
	MarketStatePrePre  = "PREPRE"
	MarketStateRegular = "REGULAR"
	MarketStatePost    = "POST"
	MarketStateClosed  = "CLOSED"
)

type financeResponse struct {
	QuoteSummary quoteSummary `json:"quoteSummary"`
}

type quoteSummary struct {
	Result []quoteSummaryResult `json:"result"`
	Error  string               `json:"error"`
}

type quoteSummaryResult struct {
	Price quoteSummaryResultPrice `json:"price"`
}

type quoteSummaryResultPrice struct {
	PreMarketPrice     marketPrice `json:"preMarketPrice"`
	RegularMarketPrice marketPrice `json:"regularMarketPrice"`
	PostMarketPrice    marketPrice `json:"postMarketPrice"`
	MarketState        string      `json:"marketState"`
}

type marketPrice struct {
	Raw decimal.Decimal `json:"raw"`
	Fmt string          `json:"fmt"`
}

type YahooClient struct {
	client *http.Client
}

func NewYahooClient() YahooClient {
	return YahooClient{
		client: &http.Client{},
	}
}

func (c YahooClient) GetYahooPrice(symbol string) (decimal.Decimal, error) {
	url := yahooFinancePriceURL + "/" + symbol + "?modules=price"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return decimal.Decimal{}, err
	}

	request.Header.Set("Content-type", "application/json")

	resp, err := c.client.Do(request)
	if err != nil {
		return decimal.Decimal{}, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return decimal.Decimal{}, err
	}

	if err = resp.Body.Close(); err != nil {
		return decimal.Decimal{}, err
	}

	var res financeResponse
	if err := json.Unmarshal(respBody, &res); err != nil {
		return decimal.Decimal{}, err
	}

	var actualPrice decimal.Decimal
	switch res.QuoteSummary.Result[0].Price.MarketState {
	case MarketStatePre:
		actualPrice = res.QuoteSummary.Result[0].Price.PreMarketPrice.Raw
	case MarketStatePrePre, MarketStateRegular:
		actualPrice = res.QuoteSummary.Result[0].Price.RegularMarketPrice.Raw
	case MarketStatePost, MarketStateClosed:
		actualPrice = res.QuoteSummary.Result[0].Price.PostMarketPrice.Raw
	}

	return actualPrice, nil
}
