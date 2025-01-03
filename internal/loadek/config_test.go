package loadek

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg := NewConfig()
	assert.Equal(t, cfg.CpuLoadMi, 100)
	assert.Equal(t, cfg.MemLoadMb, 100)

	cfg = NewConfig(
		WithCpuLoadMi(400),
		WithMemLoadMb(400),
	)

	assert.Equal(t, cfg.CpuLoadMi, 400)
	assert.Equal(t, cfg.MemLoadMb, 400)
}
