package loadek

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	app := NewApp()
	assert.Nil(t, app.Config)
	assert.Nil(t, app.Logger)

	cfg, _ := NewConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	app = NewApp(WithConfig(cfg), WithLogger(logger))

	assert.NotNil(t, app.Config)
	assert.NotNil(t, app.Logger)
}
