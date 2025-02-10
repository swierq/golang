package main

import (
	"net/http"

	"github.com/swierq/golang/internal/uihtmx"
	"github.com/swierq/golang/pkg/webapp"
)

// TODO: clean this
func (app *loadekApp) configHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	//htmx request
	if r.Header.Get("HX-Request") == "true" {
		//w.Header().Set("HX-Trigger", "config")
		err = uihtmx.RenderPartial(w, ConfigPartial(app.loadek.Config), "Dashboard", "Description")
		if err != nil {
			app.webapp.LogError(r, err)
		}
		return
	}

	//api request
	err = app.webapp.WriteJSON(w, http.StatusOK, webapp.Envelope{"config": app.loadek.Config}, nil)
	if err != nil {
		app.webapp.LogError(r, err)
	}
}

func (app *loadekApp) uiHandler(w http.ResponseWriter, r *http.Request) {
	_ = uihtmx.RenderPage(w, DashboardPage(), "UI", "Description")
}
