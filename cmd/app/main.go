package main

import (
	"github.com/lastdoctor/emma-app-go/internal/http"
	"github.com/lastdoctor/emma-app-go/internal/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	// Init Logger
	zap.ReplaceGlobals(logger.InitLog())

	// Init HTTP
	go func() {
		_, err := http.NewServer()
		if err != nil {
			zap.L().Error("fa")
			return
		}
	}()
	//pgCfg, err2 := InitRepositories()
	//if err2 != nil {
	//	return
	//}
	//log.Println(*pgCfg)

	//Init Configuration
	//const path = "./"
	//cfg, err := ConfigInit(path)
	//if err != nil {
	//	zap.L().Error("Init configuration is failed")
	//	return
	//}
	//print(cfg)
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
