package main

import (
	//"context"
	//"github.com/lastdoctor/emma-app-go/internal/app"
	//"github.com/lastdoctor/emma-app-go/internal/http"
	//"github.com/lastdoctor/emma-app-go/internal/pkg/database/postgres"
	//"github.com/lastdoctor/emma-app-go/internal/repository"
	//"log"
	//"os"
	//"os/signal"
	//"syscall"
	//"time"
	"go.uber.org/zap"
	"log"
)

func main() {
	// Init logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
		return
	}
	cfg, err := Init(configsDir)
	if err != nil {
		logger.Error("Configuration is not initialize", zap.Error(err))
		return
	}

	//
	//db, err := postgres.NewClient(cfg)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer db.Close()
	//logger.Info("Database connection established")
	//
	//// Services, Repos & API routes handler
	//repos := repository.NewRepositories(db)
	//
	//routes := http.NewRoutes(services)
	//
	//// HTTP Server
	//srv := http3.NewServer(cfg, routes.Init(cfg))
	//
	//go func() {
	//	err := srv.Run()
	//	if !errors.Is(err, http2.ErrServerClosed) {
	//		logger.Errorf("error occurred while running http server: %s\n", err.Error())
	//	}
	//}()
	//logger.Info("Server started")
	//
	//// Graceful Shutdown
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
	//const timeout = 5 * time.Second
	//ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	//defer shutdown()
	//
	//if err = srv.Stop(ctx); err != nil {
	//	logger.Errorf("failed to stop server: %v", err)
	//}
	//if err = db.Close(); err != nil {
	//	logger.Errorf("failed to disconnect: %v", err)
	//}
}
