package webapp

import (
	"fmt"
	"log/slog"
)

type Config struct {
	Port     uint16
	LogLevel string
}

type AppOption func(*Config)

func NewConfig(options ...AppOption) *Config {
	cfg := &Config{
		Port:     8080,
		LogLevel: "error",
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

func (cfg Config) GetSlogLevel() (slog.Level, error) {
	switch cfg.LogLevel {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelError, fmt.Errorf("Log level %s not found, using default log level - error", cfg.LogLevel)
	}
}
