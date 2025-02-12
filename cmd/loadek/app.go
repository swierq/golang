package main

import (
	"fmt"
	"net/http"

	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/internal/uihtmx/ui"
	"github.com/swierq/golang/pkg/webapp"
)

type loadekApp struct {
	loadek *loadek.App
	webapp *webapp.App
}

func (app *loadekApp) setupRouter() {
	app.webapp.Router.HandleFunc("GET /{$}", app.uiHandler)
	app.webapp.Router.HandleFunc("GET /api/config", app.configHandler)
	app.webapp.Router.Handle("GET /assets/", http.FileServerFS(ui.Assets))
}

type AppOption func(*loadekApp)

func newApp(options ...AppOption) (*loadekApp, error) {
	app := &loadekApp{}

	for _, option := range options {
		option(app)
	}
	if app.loadek == nil || app.webapp == nil {
		return nil, fmt.Errorf("loadek and webapp must be provided")
	}
	app.setupRouter()
	return app, nil
}

func withLoadek(loadek *loadek.App) AppOption {
	return func(app *loadekApp) {
		app.loadek = loadek
	}
}

func withWebApp(webapp *webapp.App) AppOption {
	return func(app *loadekApp) {
		app.webapp = webapp
	}
}
