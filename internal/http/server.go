package http

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HTTPConfig struct {
	Host               string        `mapstructure:"HOST"`
	Port               string        `mapstructure:"PORT"`
	ReadTimeout        time.Duration `mapstructure:"READ_TIMEOUT"`
	WriteTimeout       time.Duration `mapstructure:"WRITE_TIMEOUT"`
	MaxHeaderMegabytes int           `mapstructure:"MAX_HEADER_BYTES"`
}

func httpConfig() (*HTTPConfig, error) {
	var cfg HTTPConfig
	const path = "."
	viper.AddConfigPath(path)
	viper.SetConfigName("development")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Error("reading .env is failed", zap.Error(err))
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	return &cfg, nil
}

type Server struct {
	httpServer *http.Server
}

func NewServer() (*Server, error) {
	// Init routes
	//router := httprouter.New()
	// router.NotFound =
	// router.MethodNotAllowed =

	//router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHander)
	//router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUserHandler)
	//router.HandlerFunc(http.MethodPost, "/v1/users", app.postUserHandler)
	// HTTP config
	cgf, err := httpConfig()
	if err != nil {
		return nil, err
	}
	if err != nil {
		zap.L().Error("HTTP config is failed", zap.Error(err))
		return nil, err
	}
	port := cgf.Port
	host := cgf.Host
	read := cgf.ReadTimeout
	write := cgf.WriteTimeout
	maxheader := cgf.MaxHeaderMegabytes
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", host, port),
			//Handler:        handler,
			IdleTimeout:    time.Minute,
			ReadTimeout:    read,
			WriteTimeout:   write,
			MaxHeaderBytes: maxheader,
		},
	}, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
