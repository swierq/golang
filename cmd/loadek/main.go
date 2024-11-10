package main

import (
	"net/http"
	"strconv"

	"github.com/posener/cmd"
	"github.com/rs/zerolog"
	"github.com/swierq/golang/internal/loadek"
	"github.com/swierq/golang/pkg/webapp"
)

type loadekApp struct {
	webapp *webapp.App
}

func main() {
	root := cmd.New()
	port := root.Int("port", 8080, "Listen Port")
	debug := root.Bool("debug", false, "Debug logging")
	_ = root.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	config := webapp.NewConfig(webapp.WithPort(uint16(*port)))
	app := &loadekApp{webapp: webapp.NewApp(config)}

	app.webapp.Router.HandleFunc("GET /cpuload/{number}", app.cpuLoad)
	err := app.webapp.Serve()
	if err != nil {
		panic(err)
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
	result, err := loadek.CPULoad(number)

	if err != nil {
		app.webapp.ServerErrorResponse(w, r, err)
		return
	}
	err = app.webapp.WriteJSON(w, http.StatusOK, webapp.Envelope{"result": result}, nil)
	if err != nil {
		app.webapp.LogError(r, err)
	}
}
