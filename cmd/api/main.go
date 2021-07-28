package main

import (
	"context"
	"database/sql"
	"github.com/lastdoctor/emma-app-go/internal/data"
	"github.com/lastdoctor/emma-app-go/internal/jsonlog"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	model  data.Model
	wg     sync.WaitGroup
}

func main() {

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: cfg,
		logger: logger,
		model:  data.NewModel(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
