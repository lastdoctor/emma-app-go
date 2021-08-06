package handlers

import (
	"net/http"
)

func (app *main.application) healthcheckHander(w http.ResponseWriter, r *http.Request) {
	envo := envelope{
		"isOK": true,
	}
	err := app.writeJSON(w, http.StatusOK, envo, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}