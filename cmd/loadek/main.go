package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/posener/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/swierq/golang/internal/loadek"
)

func main() {
	root := cmd.New()
	port := root.String("port", "8080", "Listen Port")
	debug := root.Bool("debug", false, "Debug logging")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	router := mux.NewRouter()
	router.HandleFunc("/cpuload/{number}", cpuLoad)
	log.Info().Msgf("Starting Server on port: %s", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), router)
	if err != nil {
		log.Error().Msg("Something went wrong. Exiting.")
		panic(err)
	}
}

func cpuLoad(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.ParseInt(vars["number"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	result, err := loadek.CPULoad(number)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, err = fmt.Fprint(w, result)
	if err != nil {
		panic(err)
	}
}
