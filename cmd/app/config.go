package main

import (
	"flag"
	"github.com/spf13/viper"
	"time"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		HTTP        HTTPConfig
		Limiter     LimiterConfig
	}

	PostgresConfig struct {
		HOST      string
		User      string
		Password  string
		Database  string
		Port      int
		OpenConns int
		IdleConns int
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
)

func Init(configsDir string) (*Config, error) {
	var cfg = populateDefaults()
	return cfg, nil

	// if err := parseConfigFile(configsDir, os.Getenv("APP_ENV")); err != nil {
	//	return nil, err
	// }
	//
	// var cfg Config
	// if err := unmarshal(&cfg); err != nil {
	//	return nil, err
	// }

	// setFromEnv(&cfg)
}

func populateDefaults() *Config {
	var cfg Config
	flag.StringVar(&cfg.HTTP.Port, "port", "4000", "API server port")
	flag.StringVar(&cfg.Environment, "environment", "development", "Environment (development|staging|production)")
	// flag.StringVar(&cfg.db.dsn, "db-dsn", "host=localhost port=5432 user=postgres dbname=postgres password=root sslmode=disable", "Postgres DSN")
	flag.StringVar(&cfg.Postgres.HOST, "db-host", "localhost", "PostgreSQL host")
	flag.StringVar(&cfg.Postgres.Database, "db-database", "postgres", "PostgreSQL database")
	flag.StringVar(&cfg.Postgres.User, "db-user", "postgres", "PostgreSQL user")
	flag.StringVar(&cfg.Postgres.Password, "db-password", "root", "PostgreSQL password")
	flag.IntVar(&cfg.Postgres.Port, "db-port", 5432, "PostgreSQL Port")
	flag.IntVar(&cfg.Postgres.OpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.Postgres.IdleConns, "db-max-open-idle", 25, "PostgreSQL max open idle")
	flag.IntVar(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.DurationVar(&cfg.Limiter.TTL, "limiter-ttl", 3*time.Second, "Rate limiter maximum burst")
	flag.Parse()

	viper.SetDefault("db.host", cfg.Postgres.HOST)
	viper.SetDefault("db.user", cfg.Postgres.User)
	viper.SetDefault("db.password", cfg.Postgres.Password)
	viper.SetDefault("db.database", cfg.Postgres.Database)
	viper.SetDefault("db.port", cfg.Postgres.Port)
	viper.SetDefault("db.openconns", cfg.Postgres.OpenConns)
	viper.SetDefault("db.idleconns", cfg.Postgres.IdleConns)
	viper.SetDefault("http.port", cfg.HTTP.Port)
	viper.SetDefault("http.max_header_megabytes", cfg.HTTP.MaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", cfg.HTTP.ReadTimeout)
	viper.SetDefault("http.timeouts.write", cfg.HTTP.WriteTimeout)
	viper.SetDefault("limiter.rps", cfg.Limiter.RPS)
	viper.SetDefault("limiter.burst", cfg.Limiter.Burst)
	viper.SetDefault("limiter.ttl", cfg.Limiter.TTL)
	return &cfg
}

//func unmarshal(cfg *Config) error {
//	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
//		return err
//	}
//
//	if err := viper.UnmarshalKey("limiter", &cfg.Limiter); err != nil {
//		return err
//	}
//
//	return nil
//}

//func setFromEnv(cfg *Config) {
//	cfg.Postgres.HOST = os.Getenv("POSTGRES_URI")
//	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
//	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
//	cfg.Postgres.Database = os.Getenv("POSTGRES_DB")
//
//	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
//}
