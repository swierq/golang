package loadek

import "log/slog"

type App struct {
	Config *Config
	Logger *slog.Logger
}

type AppOption func(*App)

func NewApp(options ...AppOption) *App {
	app := &App{}
	for _, option := range options {
		option(app)
	}
	return app
}

func WithConfig(config *Config) AppOption {
	return func(a *App) {
		a.Config = config
	}
}

func WithLogger(logger *slog.Logger) AppOption {
	return func(a *App) {
		a.Logger = logger
	}
}
