package webapp

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg.Port, uint16(8080))

	logLevel, err := cfg.GetSlogLevel()
	assert.Equal(t, logLevel, slog.LevelError)
	assert.Nil(t, err)

	cfg = NewConfig(
		WithPort(uint16(4000)),
		WithLogLevel("debug"),
	)
	assert.Equal(t, cfg.Port, uint16(4000))

	logLevel, err = cfg.GetSlogLevel()
	assert.Equal(t, logLevel, slog.LevelDebug)
	assert.Nil(t, err)

	cfg = NewConfig(
		WithLogLevel("warn"),
	)
	logLevel, err = cfg.GetSlogLevel()
	assert.Equal(t, logLevel, slog.LevelWarn)
	assert.Nil(t, err)

	cfg = NewConfig(
		WithLogLevel("info"),
	)
	logLevel, err = cfg.GetSlogLevel()
	assert.Equal(t, logLevel, slog.LevelInfo)
	assert.Nil(t, err)

	cfg = NewConfig(
		WithLogLevel("badValue"),
	)
	logLevel, err = cfg.GetSlogLevel()
	assert.Equal(t, logLevel, slog.LevelError)
	assert.NotNil(t, err)

}
