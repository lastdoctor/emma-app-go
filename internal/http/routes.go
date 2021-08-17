package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)
	return router
}
