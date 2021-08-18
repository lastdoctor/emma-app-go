package pg

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/lastdoctor/emma-app-go/internal/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

// Timeout is Postgres timeout
const Timeout = 5

type PostgresConfig struct {
	PgHost    string
	User      string
	Password  string
	Database  string
	PgPort    int
	OpenConns int
	IdleConns int
}

func New(ctx context.Context, cfg PostgresConfig) (*pg.DB, error) {
	// Connect to Postgres database
	pgDB, err := Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "pg.Dial failed")
	}

	// Run Postgres migrations
	// TODO`implement`

	var pg *pg.DB

	// Initialize the Postgres repositories
	if pgDB != nil {
		pg = pgDB
		go KeepAlivePg(pg, cfg)
		// implement the function
		// TODO `implement`
		//User = pg.NewUserRepo(pgDB)
		//Merchant = pg.NewMerchantRepo(pgDB)
		//Transaction = pg.TransactionRepo(pgDB)
	}
	return pg, nil
}

// KeepAlivePollPeriod is a Pg/MySQL keepalive check time period
const KeepAlivePollPeriod = 3

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func KeepAlivePg(pg *pg.DB, cfg PostgresConfig) {
	var err error
	for {
		// Check if PostgreSQL is alive every 3 seconds
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false
		if pg == nil {
			lostConnect = true
		} else if _, err = pg.Exec("SELECT 1"); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		logger.Logger().Debug("[store.KeepAlivePg] Lost PostgreSQL connection. Restoring...")
		_, err = Dial(cfg)
		if err != nil {
			logger.Logger().Error("Failed to connect to Postgres", zap.Error(err))
			continue
		}
		logger.Logger().Debug("[store.KeepAlivePg] PostgreSQL reconnected")
	}
}

// Dial creates new database connection to postgres
func Dial(cfg PostgresConfig) (*pg.DB, error) {
	pgURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.PgHost, cfg.PgPort, cfg.Database)
	pgOpt, err := pg.ParseURL(pgURL)
	pgDB := pg.Connect(pgOpt)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return pgDB, nil
}
