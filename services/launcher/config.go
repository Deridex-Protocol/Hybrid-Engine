package launcher

type Config struct {
	DatabaseURL       string `env:"DERIDEX_DATABASE_URL"`
	EthereumRpcURL    string `env:"DERIDEX_ETHEREUM_RPC_URL"`
	ProxyAddress      string `env:"DERIDEX_PROXY_ADDRESS"`
	RelayerPrivateKey string `env:"DERIDEX_RELAYER_PK"`
	GasLimit          uint64 `env:"DERIDEX_GAS_LIMIT"`
	ManualGasPrice    int64  `env:"DERIDEX_MANUAL_GAS_PRICE"`
}
