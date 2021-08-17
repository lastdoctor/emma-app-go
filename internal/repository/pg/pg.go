package pg

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"time"
)

// Timeout is Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	cfg, _ := pgConfig()
	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.HOST, cfg.Port, cfg.Database)
	pgOpt, err := pg.ParseURL(pgURL)
	pgDB := pg.Connect(pgOpt)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return &DB{pgDB}, nil
}
