package http

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

package http

import (
"context"
"github.com/julienschmidt/httprouter"
"github.com/lastdoctor/emma-app-go/cmd/app"
"net/http"
"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *main.Config, handler http.Handler) *Server {
	router := httprouter.New()
	// router.NotFound =
	// router.MethodNotAllowed =

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHander)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.postUserHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.HTTP.Port,
			Handler:        handler,
			IdleTimeout:    time.Minute,
			ReadTimeout:    cfg.HTTP.ReadTimeout,
			WriteTimeout:   cfg.HTTP.WriteTimeout,
			MaxHeaderBytes: cfg.HTTP.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
