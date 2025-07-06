package main

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swierq/golang/internal/uihtmx/ui"
)

type config struct {
	Port int `json:"port" yaml:"port"`
}

type app struct {
	config *config
	e      *echo.Echo
}

type configOption func(*config) error

func newConfig(options ...configOption) (*config, error) {
	cfg := &config{
		Port: 1323, // Default port
	}

	for _, option := range options {
		err := option(cfg)
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}
func withConfigPort(port int) configOption {
	return func(cfg *config) error {
		cfg.Port = port
		return nil
	}
}

type appOption func(*app) error

func newApp(options ...appOption) (*app, error) {
	app := &app{}
	err := app.initEcho()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize echo: %w", err)
	}

	for _, option := range options {
		err := option(app)
		if err != nil {
			return nil, err
		}
	}
	if app.config == nil {
		return nil, fmt.Errorf("config must be provided")
	}
	return app, nil
}

func withConfig(cfg *config) appOption {
	return func(app *app) error {
		if cfg == nil {
			return fmt.Errorf("config cannot be nil")
		}
		app.config = cfg
		return nil
	}
}

func (app *app) initEcho() error {
	assets, _ := fs.Sub(ui.Assets, "assets")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/headers", app.printHeaders)
	e.GET("/setcookie", app.writeCookie)
	e.GET("/cookies", app.printCookies)
	e.StaticFS("/assets/", assets)
	app.e = e
	return nil
}

func (app *app) Start() error {
	app.e.Logger.Fatal(app.e.Start(fmt.Sprintf(":%d", app.config.Port)))
	return nil
}
