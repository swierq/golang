package webapp

import (
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	Config *Config
	Logger *slog.Logger
	Router *http.ServeMux
}

func NewApp(config *Config) *App {
	slogLevel, err := config.GetSlogLevel()
	app := &App{
		Config: config,
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel,
		})),
		Router: http.NewServeMux(),
	}
	app.Router.HandleFunc("GET /health", app.HealthcheckHandler)
	if err != nil {
		app.Logger.Warn("Default log level used", "error", err.Error())
	}
	return app
}

func (app *App) GetRouter() http.Handler {
	return app.recoverPanic(secureHeaders(app.Router))
}
