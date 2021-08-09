package pg

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
	"log"
	"time"
)

// Timeout is Postgres timeout
const Timeout = 5

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

type PostgresConfig struct {
	HOST      string `mapstructure:"PG_HOST"`
	User      string `mapstructure:"PG_USER"`
	Password  string `mapstructure:"PG_PASSWORD"`
	Database  string `mapstructure:"PG_DATABASE"`
	Port      int    `mapstructure:"PG_PORT"`
	OpenConns int    `mapstructure:"PG_CONNS"`
	IdleConns int    `mapstructure:"PG_IDLE"`
}

// Dial creates new database connection to postgres
func Dial() (*DB, error) {
	const pi = 3.1
	log.Println(pi)
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

func pgConfig() (*PostgresConfig, error) {
	var pgCfg PostgresConfig
	const path = "."
	viper.AddConfigPath(path)
	viper.SetConfigName("development")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&pgCfg)
	return &pgCfg, nil
}
