package main

import (
	"context"
	"github.com/lastdoctor/emma-app-go/internal/http"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"github.com/lastdoctor/emma-app-go/internal/pg"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := run(); err != nil {
		logger.Logger().Fatal("Cannot start Emma App", zap.Error(err))
	}
}

func run() error {
	// Create context
	ctx := context.Background()

	// Initialize configuration
	cfg, err := GetConfig()
	if err != nil {
		logger.Logger().Error("Failed to get configuration", zap.Error(err))
		return err
	}

	// Initialize postgres and repository
	pgCfd := pg.PostgresConfig{
		PgHost:    cfg.PgHost,
		User:      cfg.User,
		Password:  cfg.Password,
		Database:  cfg.Database,
		PgPort:    cfg.PgPort,
		OpenConns: cfg.OpenConns,
		IdleConns: cfg.IdleConns,
	}
	pgDB, err := pg.New(ctx, pgCfd)
	if err != nil {
		logger.Logger().Error("Failed to initialize the store", zap.Error(err))
		return err
	}
	log.Println(pgDB)
	// TODO `Initialize and implement service`

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
	return nil
}
