package loadek

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cfg, err := NewConfig()
	assert.Nil(t, err)
	assert.Equal(t, cfg.CpuLoadMi, 100)
	assert.Equal(t, cfg.MemLoadMb, 100)

	cfg, err = NewConfig(
		WithCpuLoadMi(400),
		WithMemLoadMb(400),
	)
	assert.Nil(t, err)
	assert.Equal(t, cfg.CpuLoadMi, 400)
	assert.Equal(t, cfg.MemLoadMb, 400)

	cfg, err = NewConfig(
		WithCpuLoadMi(-1),
		WithMemLoadMb(400),
	)
	assert.Nil(t, cfg)
	assert.NotNil(t, err)

	cfg, err = NewConfig(
		WithCpuLoadMi(100),
		WithMemLoadMb(-1),
	)
	assert.Nil(t, cfg)
	assert.NotNil(t, err)
}
