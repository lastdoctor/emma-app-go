package http

//
//import (
//	"net/http"
//)
//
//func (app *application) logError(r *http.Request, err error) {
//	app.logger.PrintError(err, map[string]string{
//		"request_method": r.Method,
//		"request_url":    r.URL.String(),
//	})
//}
//func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
//	env := envelope{"error": message}
//	err := app.writeJSON(w, status, env, nil)
//	if err != nil {
//		app.logError(r, err)
//		w.WriteHeader(status)
//	}
//}
//
//func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
//	app.logError(r, err)
//	message := "the server encountered a problem and could not process your request"
//	app.errorResponse(w, r, http.StatusInternalServerError, message)
//}
//
//func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
//	message := "the request resource not found"
//	app.errorResponse(w, r, http.StatusNotFound, message)
//}
//
//func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
//	app.logError(r, err)
//	message := "the request is invalid"
//	app.errorResponse(w, r, http.StatusBadRequest, message)
//}
//
//func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
//	message := "the request method is not allowed"
//	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
//}
