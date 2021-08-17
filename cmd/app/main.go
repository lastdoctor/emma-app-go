package main

import (
	"github.com/lastdoctor/emma-app-go/internal/http"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		logger.Logger().Fatal("Cannot start EmmaApp", zap.Error(err))
	}
}

func run() error {
	//ctx := context.contextBackground()

	// Initialize configuration
	cfg, err := GetConfig()
	if err != nil {
		logger.Logger().Error("Failed to get configuration", zap.Error(err))
		return err
	}
	// Starting API server
	httpCfg := http.HTTPConfig{
		PORT:         cfg.PORT,
		RPS:          cfg.RPS,
		BURST:        cfg.BURST,
		TTL:          cfg.TTL,
		HOST:         cfg.HOST,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	err = http.Serve(httpCfg)
	if err != nil {
		logger.Logger().Error("Failed to start HTTP server", zap.Error(err))
		return err
	}

	//Init repository repository
	//repository, err := repository.New(ctx)
	//if err != nil {
	//	logger.Logger().Error(err)
	//}

	//dial, err := pg.Dial()
	//if err != nil {
	//	logger.Logger().Error("Failed to connect to database", zap.Error(err))
	//	return err
	//}
	//logger.Logger().Info(dial.String())
	return nil
}
