package launcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

const getGasPriceURL = "https://ethgasstation.info/json/ethgasAPI.json"

type GasPriceRespBody struct {
	Fast    decimal.Decimal `json:"fast"`
	Average decimal.Decimal `json:"average"`
}

func GasPriceInWei() (decimal.Decimal, error) {
	resp, err := http.Get(getGasPriceURL)
	if err != nil {
		return decimal.Decimal{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return decimal.Decimal{}, err
	}

	var gasPriceResp GasPriceRespBody
	if err = json.Unmarshal(body, &gasPriceResp); err != nil {
		return decimal.Decimal{}, err
	}

	// returned value from gasStation api is in 0.1 Gwei
	return gasPriceResp.Fast.Mul(decimal.New(1, 8)), nil
}
