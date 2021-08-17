package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HTTPConfig struct {
	PORT         int
	RPS          int
	BURST        int
	TTL          int
	HOST         string
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

func Serve(cfg HTTPConfig) error {
	srv := http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.HOST, cfg.PORT),
		Handler:        routes(),
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    time.Duration(cfg.IdleTimeout),
		ReadTimeout:    time.Duration(cfg.ReadTimeout),
		WriteTimeout:   time.Duration(cfg.WriteTimeout),
	}

	// Create a shutdown channel to receive any errors
	// returned by the graceful Shutdown function
	shutdownError := make(chan error)

	// Start a background goroutine HTTP
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		logger.Logger().Info(fmt.Sprintf("Caught signal shutdown: %v", s))

		// Create a context with 5-second timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown Error channel if returns any error
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		logger.Logger().Info(fmt.Sprintf("Completed graceful shutdown Addr %s", srv.Addr))
		shutdownError <- nil
	}()
	// Start the API server
	logger.Logger().Info(fmt.Sprintf("Starting API server address: %s", srv.Addr))
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	err = <-shutdownError
	if err != nil {
		return err
	}
	// Graceful shutdown API server
	logger.Logger().Info(fmt.Sprintf("Shutdown API server address: %s", srv.Addr))

	return nil
}
