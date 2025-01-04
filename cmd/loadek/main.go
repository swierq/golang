package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/posener/cmd"
	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/pkg/webapp"
)

func main() {
	root := cmd.New()
	port := root.Int("port", 8080, "Listen Port")
	cpuMi := root.Int("cpumi", 100, "Cpu milocores")
	memMb := root.Int("memmb", 20, "Memory mb")
	_ = root.Parse()

	//TODO: should return error also
	webConfig := webapp.NewConfig(webapp.WithPort(uint16(*port)), webapp.WithLogLevel("info"))
	//TODO: should return error also
	webapp := webapp.NewApp(webConfig)

	//TODO: should return error also
	loadekConfig := loadek.NewConfig(loadek.WithCpuLoadMi(*cpuMi), loadek.WithMemLoadMb(*memMb))
	//TODO: should return error also
	loadek := loadek.NewApp(loadek.WithConfig(loadekConfig))

	app, err := newApp(withLoadek(loadek), withWebApp(webapp))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go waitOnSignal(cancel)
	err = app.webapp.Serve(ctx)
	if err != nil {
		app.webapp.Logger.Error("Webserver problem", "error", err)
	}
}

func waitOnSignal(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
}
