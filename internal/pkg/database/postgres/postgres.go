package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lastdoctor/emma-app-go/internal/config"
	"time"
)

func NewClient(cfg *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.HOST, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.Postgres.OpenConns)
	db.SetMaxIdleConns(cfg.Postgres.IdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
