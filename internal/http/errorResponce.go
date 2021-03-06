package http

import (
	"fmt"
	"github.com/lastdoctor/emma-app-go/internal/controller"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

func serverErrorResponse(w http.ResponseWriter, r *http.Request) {
	message := "The server error occurred"
	err := controller.WriteJSON(w, http.StatusInternalServerError, message, nil)
	if err != nil {
		if err != nil {
			logger.Logger().Error(fmt.Sprintf("JSON encoding failed: %v, requested by %s", err, r.RequestURI), zap.Error(err))
		}
	}
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested method is not allowed"
	err := controller.WriteJSON(w, http.StatusMethodNotAllowed, message, nil)
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("JSON encoding failed: %v, requested by %s", err, r.RequestURI), zap.Error(err))
	}
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	err := controller.WriteJSON(w, http.StatusNotFound, message, nil)
	if err != nil {
		logger.Logger().Error(fmt.Sprintf("JSON encoding failed: %v, requested by %s", err, r.RequestURI), zap.Error(err))
	}
}
