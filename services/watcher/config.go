package watcher

type Config struct {
	DatabaseURL    string `env:"DERIDEX_DATABASE_URL"`
	RedisURL       string `env:"DERIDEX_REDIS_URL"`
	EthereumRpcURL string `env:"DERIDEX_ETHEREUM_RPC_URL"`
}
