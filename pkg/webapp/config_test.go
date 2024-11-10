package webapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg.Port, uint16(8080))
	assert.Equal(t, cfg.LogLevel, "info")

	cfg = NewConfig(
		WithPort(uint16(4000)),
		WithLogLevel("debug"),
	)
	assert.Equal(t, cfg.Port, uint16(4000))
	assert.Equal(t, cfg.LogLevel, "debug")

}
