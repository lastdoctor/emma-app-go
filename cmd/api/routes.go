package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHander)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.postUserHandler)
	return router
}
