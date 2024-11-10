package webapp

type Config struct {
	Port     uint16
	LogLevel string
}

type AppOption func(*Config)

func NewConfig(options ...AppOption) *Config {
	cfg := &Config{
		Port:     8080,
		LogLevel: "info",
	}

	for _, option := range options {
		option(cfg)
	}
	return cfg
}

func WithPort(port uint16) AppOption {
	return func(c *Config) {
		c.Port = port
	}
}

func WithLogLevel(level string) AppOption {
	return func(c *Config) {
		c.LogLevel = level
	}
}
