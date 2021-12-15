package websocket

type Config struct {
	RedisURL string   `env:"DERIDEX_REDIS_URL"`
	ApiURL   string   `env:"DERIDEX_API_URL"`
	Channels []string `env:"DERIDEX_CHANNELS"`
}
