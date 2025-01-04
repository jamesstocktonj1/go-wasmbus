package wasmbus

// Config
type Config struct {
	Name    string
	NatsUrl string
}

func defaultConfig(cfg *Config) (*Config, error) {
	if len(cfg.Name) < 1 {
		cfg.Name = "go-wasmbus"
	}

	if len(cfg.NatsUrl) < 1 {
		cfg.NatsUrl = "localhost:4222"
	}

	return cfg, nil
}
