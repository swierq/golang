package main

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/labstack/echo/v4"
	"github.com/swierq/golang/internal/uihtmx/ui"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

type config struct {
	Port         int    `json:"port" yaml:"port"`
	ClientID     string `json:"clientID" yaml:"clientID"`
	ClientSecret string `json:"clientSecret" yaml:"clientSecret"`
	RedirectURL  string `json:"redirectURL" yaml:"redirectURL"`
	TenantID     string `json:"tenantID" yaml:"tenantID"`
	AppScope     string `json:"appScope" yaml:"appScope"`
}

type app struct {
	config *config
	e      *echo.Echo
	entra  *oauth2.Config
	jwks   *keyfunc.JWKS
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

func withConfigClientID(clientID string) configOption {
	return func(cfg *config) error {
		cfg.ClientID = clientID
		return nil
	}
}

func withConfigClientSecret(clientSecret string) configOption {
	return func(cfg *config) error {
		cfg.ClientSecret = clientSecret
		return nil
	}
}

func withConfigTenantID(tenantID string) configOption {
	return func(cfg *config) error {
		cfg.TenantID = tenantID
		return nil
	}
}

func withConfigRedirectURL(redirectURL string) configOption {
	return func(cfg *config) error {
		cfg.RedirectURL = redirectURL
		return nil
	}
}

func withConfigAppScope(appScope string) configOption {
	return func(cfg *config) error {
		cfg.AppScope = appScope
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

	// Set default values for EntraID OpenID authentication
	if app.config.ClientID == "" {
		return nil, fmt.Errorf("clientID must be provided")
	}
	if app.config.ClientSecret == "" {
		return nil, fmt.Errorf("clientSecret must be provided")
	}
	if app.config.RedirectURL == "" {
		return nil, fmt.Errorf("redirectURL must be provided")
	}

	if app.config.TenantID == "" {
		return nil, fmt.Errorf("tenantID must be provided")
	}

	if app.config.AppScope == "" {
		return nil, fmt.Errorf("appScope must be provided")
	}

	app.entra = &oauth2.Config{
		ClientID:     app.config.ClientID,
		ClientSecret: app.config.ClientSecret,
		RedirectURL:  app.config.RedirectURL,
		Endpoint:     microsoft.AzureADEndpoint(app.config.TenantID),
		Scopes:       []string{"openid", "profile", "email", app.config.AppScope},
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

	e.GET("/", app.homeHandler)
	e.GET("/login", app.loginHandler)
	e.GET("/token", app.tokenHandler)

	e.GET("/callback", app.callbackHandler)

	e.GET("/logout", app.logoutHandler)
	e.GET("/headers", app.printHeaders)
	e.GET("/setcookie", app.writeCookie)
	e.GET("/cookies", app.printCookies)
	e.StaticFS("/assets/", assets)
	app.e = e
	return nil
}

func (app *app) Start(ctx context.Context) error {

	//initialize JWKS (JSON Web Key Set) for token validation
	jwksOptions := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			app.e.Logger.Error("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwksURL := fmt.Sprintf("https://login.microsoftonline.com/%s/discovery/v2.0/keys", app.config.TenantID)
	jwks, err := keyfunc.Get(jwksURL, jwksOptions)
	if err != nil {
		app.e.Logger.Error("Failed to create JWKS from resource at the given URL: %s", err.Error())
		return err
	}
	app.jwks = jwks
	defer app.jwks.EndBackground()

	//Start the Echo server
	go func() {
		if err := app.e.Start(fmt.Sprintf(":%d", app.config.Port)); err != nil && err != http.ErrServerClosed {
			app.e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = app.e.Shutdown(ctx)
	if err != nil {
		app.e.Logger.Fatal(err)
	}
	return err
}
