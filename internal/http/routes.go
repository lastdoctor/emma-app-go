package http

import (
	"expvar"
	"github.com/julienschmidt/httprouter"
	"github.com/lastdoctor/emma-app-go/internal/controller"
	"net/http"
)

func routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedResponse)

	//// CreateU
	router.HandlerFunc(http.MethodPost, "/v1/users/:id", controller.GetUserById)
	router.HandlerFunc(http.MethodPost, "/v1/users", controller.CreateUser)
	router.HandlerFunc(http.MethodPut, "v1/users", controller.UpdateUser)
	router.HandlerFunc(http.MethodDelete, "v1/users/:id", controller.DeleteUser)
	//// Merchant
	//router.HandlerFunc(http.MethodPost, "/v1/merchants/:id", controller.GetMerchantById)
	//router.HandlerFunc(http.MethodPost, "/v1/merchants", controller.CreateMerchant)
	//router.HandlerFunc(http.MethodPut, "v1/merchants", controller.updateMerchant)
	//router.HandlerFunc(http.MethodDelete, "v1/merchants/:id", controller.DeleteMerchant)
	//// Transaction
	//router.HandlerFunc(http.MethodPost, "/v1/transactions/:id", controller.GetTransactionById)
	//router.HandlerFunc(http.MethodPost, "/v1/transactions", controller.CreateTransaction)
	//router.HandlerFunc(http.MethodPut, "v1/transactions", controller.UpdateTransaction)
	//router.HandlerFunc(http.MethodDelete, "v1/transactions/:id", controller.DeleteTransaction)

	// Register a new endpoint pointing to the expvar handler.
	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return router
}
