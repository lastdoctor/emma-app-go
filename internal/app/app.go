package app

import (
	"context"
	"errors"
	"github.com/lastdoctor/emma-app-go/internal/config"
	delivery "github.com/lastdoctor/emma-app-go/internal/delivery/http"
	"github.com/lastdoctor/emma-app-go/internal/pkg/database/postgres"
	"github.com/lastdoctor/emma-app-go/internal/pkg/logger"
	"github.com/lastdoctor/emma-app-go/internal/repository"
	"github.com/lastdoctor/emma-app-go/internal/server"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)

		return
	}

	db, err := postgres.NewClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	logger.Info("Database connection established")

	// Services, Repos & API routes handler
	repos := repository.NewRepositories(db)

	routes := delivery.NewRoutes(services)

	// HTTP Server
	srv := server.NewServer(cfg, routes.Init(cfg))

	go func() {
		err := srv.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Info("Server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
	if err = db.Close(); err != nil {
		logger.Errorf("failed to disconnect: %v", err)
	}
}
