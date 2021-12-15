package engine

type Config struct {
	DatabaseURL     string `env:"DERIDEX_DATABASE_URL"`
	RedisURL        string `env:"DERIDEX_REDIS_URL"`
	ExchangeAddress string `env:"DERIDEX_EXCHANGE_ADDRESS"`
	RelayerAddress  string `env:"DERIDEX_RELAYER_ADDRESS"`
}
