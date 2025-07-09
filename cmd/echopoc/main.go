package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	ca := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Value:   1323,
				Usage:   "port",
				EnvVars: []string{"ECHOPOC_PORT"},
			},
			&cli.StringFlag{
				Name:    "tenantID",
				Value:   "",
				Usage:   "tenant ID for OAuth2",
				EnvVars: []string{"ECHOPOC_TENANT_ID"},
			},
			&cli.StringFlag{
				Name:    "clientID",
				Value:   "",
				Usage:   "client ID for OAuth2",
				EnvVars: []string{"ECHOPOC_CLIENT_ID"},
			},
			&cli.StringFlag{
				Name:    "clientSecret",
				Value:   "",
				Usage:   "client Secret for OAuth2",
				EnvVars: []string{"ECHOPOC_CLIENT_SECRET"},
			},
			&cli.StringFlag{
				Name:    "redirectURL",
				Value:   "",
				Usage:   "redirect url for OAuth2",
				EnvVars: []string{"ECHOPOC_REDIRECT_URL"},
			},
			&cli.StringFlag{
				Name:    "appScope",
				Value:   "",
				Usage:   "appScope url for OAuth2",
				EnvVars: []string{"ECHOPOC_APP_SCOPE"},
			},
		},
		Action: func(cCtx *cli.Context) error {
			return run(cCtx)
		},
	}
	if err := ca.Run(os.Args); err != nil {
		slog.Error("could not start application", "error", err)
	}

}

func run(cCtx *cli.Context) error {
	cfg, err := newConfig(
		withConfigPort(cCtx.Int("port")),
		withConfigTenantID(cCtx.String("tenantID")),
		withConfigClientID(cCtx.String("clientID")),
		withConfigClientSecret(cCtx.String("clientSecret")),
		withConfigRedirectURL(cCtx.String("redirectURL")),
		withConfigAppScope(cCtx.String("appScope")),
	)

	if err != nil {
		slog.Error("failed to create config", "error", err)
		panic(err)
	}

	a, err := newApp(withConfig(cfg))
	if err != nil {
		slog.Error("failed to create app", "error", err)
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go waitOnSignal(cancel)

	return a.Start(ctx)
}

func waitOnSignal(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
}
