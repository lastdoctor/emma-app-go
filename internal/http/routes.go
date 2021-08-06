package http

//type Routes struct {
//	services *service.Services
//}
//
//func NewRoutes(services *service.Services) *Routes {
//	return &Routes{
//		services: services,
//	}
//}
//
//func (h *Routes) Init(cfg *main.Config) *httprouter.Router {
//	router := httprouter.New()
//	// router.NotFound =
//	// router.MethodNotAllowed =
//
//	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHander)
//	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUserHandler)
//	router.HandlerFunc(http.MethodPost, "/v1/users", app.postUserHandler)
//	return router
//}
