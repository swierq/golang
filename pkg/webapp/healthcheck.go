package webapp

import (
	"net/http"
)

func (app *App) HealthcheckHandler(w http.ResponseWriter, r *http.Request) {

	env := Envelope{
		"status": "available",
	}
	err := app.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.Logger.Error(err.Error())
		app.ServerErrorResponse(w, r, err)
		return
	}
}
