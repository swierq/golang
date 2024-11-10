package webapp

import (
	"log/slog"
	"net/http"
)

type App struct {
	Config *Config
	Logger *slog.Logger
	Router *http.ServeMux
}

func NewApp(config *Config) *App {
	app := &App{
		Config: config,
		Logger: slog.Default(),
		Router: http.NewServeMux(),
	}
	app.Router.HandleFunc("GET /health", app.HealthcheckHandler)
	return app
}

func (app *App) GetRouter() http.Handler {
	return app.recoverPanic(secureHeaders(app.Router))
}
