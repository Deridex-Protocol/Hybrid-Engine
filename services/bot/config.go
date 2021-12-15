package bot

type Config struct {
	ApiURL      string  `env:"DERIDEX_API_URL"`
	Address1    string  `env:"DERIDEX_BOT_ADDRESS_1"`
	PrivateKey1 string  `env:"DERIDEX_BOT_PRIVATE_KEY_1"`
	Address2    string  `env:"DERIDEX_BOT_ADDRESS_2"`
	PrivateKey2 string  `env:"DERIDEX_BOT_PRIVATE_KEY_2"`
	MarketID    string  `env:"DERIDEX_BOT_MARKET_ID"`
	MinAmount   float64 `env:"DERIDEX_BOT_MIN_AMOUNT"`
	MaxAmount   float64 `env:"DERIDEX_BOT_MAX_AMOUNT"`
	TickerTime  string  `env:"DERIDEX_BOT_TICKER_TIME"`
}
