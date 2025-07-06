package main

import (
	"log/slog"
	"os"

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
	cfg, err := newConfig(withConfigPort(cCtx.Int("port")))
	if err != nil {
		slog.Error("failed to create config", "error", err)
		panic(err)
	}

	a, err := newApp(withConfig(cfg))
	if err != nil {
		slog.Error("failed to create app", "error", err)
		panic(err)
	}
	_ = a.Start()
	return nil

}
