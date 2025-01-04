package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/posener/cmd"
	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/pkg/webapp"
)

type loadekApp struct {
	loadek *loadek.App
	webapp *webapp.App
}

func main() {
	root := cmd.New()
	port := root.Int("port", 8080, "Listen Port")
	_ = root.Parse()

	config := webapp.NewConfig(webapp.WithPort(uint16(*port)), webapp.WithLogLevel("info"))
	app := &loadekApp{webapp: webapp.NewApp(config)}

	ctx, cancel := context.WithCancel(context.Background())

	app.webapp.Router.HandleFunc("GET /cpuload/{number}", app.cpuLoad)
	go waitOnSignal(cancel)
	err := app.webapp.Serve(ctx)
	if err != nil {
		app.webapp.Logger.Error("Webserver problem", "error", err)
	}
}

func (app *loadekApp) cpuLoad(w http.ResponseWriter, r *http.Request) {
	num := r.PathValue("number")
	number64, err := strconv.ParseInt(num, 10, 0)
	number := int(number64)
	if err != nil {
		app.webapp.BadRequestResponse(w, r, err)
		return
	}
	result, err := app.loadek.CPULoad(number)

	if err != nil {
		app.webapp.ServerErrorResponse(w, r, err)
		return
	}
	err = app.webapp.WriteJSON(w, http.StatusOK, webapp.Envelope{"result": result}, nil)
	if err != nil {
		app.webapp.LogError(r, err)
	}
}

func waitOnSignal(cancel context.CancelFunc) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
}
