package api

type Config struct {
	DatabaseURL     string `env:"DERIDEX_DATABASE_URL"`
	RedisURL        string `env:"DERIDEX_REDIS_URL"`
	EthereumRpcURL  string `env:"DERIDEX_ETHEREUM_RPC_URL"`
	ChanID          int64  `env:"DERIDEX_ETHEREUM_CHAN_ID"`
	ProxyAddress    string `env:"DERIDEX_PROXY_ADDRESS"`
	ExchangeAddress string `env:"DERIDEX_EXCHANGE_ADDRESS"`
	RelayerAddress  string `env:"DERIDEX_RELAYER_ADDRESS"`
}
