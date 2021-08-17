package main

import "github.com/spf13/viper"

type Config struct {
	// Postgres configuration
	PgHost    string `mapstructure:"PG_HOST"`
	User      string `mapstructure:"PG_USER"`
	Password  string `mapstructure:"PG_PASSWORD"`
	Database  string `mapstructure:"PG_DATABASE"`
	PgPort    int    `mapstructure:"PG_PORT"`
	OpenConns int    `mapstructure:"PG_CONNS"`
	IdleConns int    `mapstructure:"PG_IDLE"`
	// HTTP configuration
	PORT         int    `mapstructure:"PORT"`
	RPS          int    `mapstructure:"RPS"`
	BURST        int    `mapstructure:"BURST"`
	TTL          int    `mapstructure:"TTL"`
	HOST         string `mapstructure:"HOST"`
	ReadTimeout  int    `mapstructure:"READ_TIMEOUT"`
	WriteTimeout int    `mapstructure:"WRITE_TIMEOUT"`
	IdleTimeout  int    `mapstructure:"IDLE"`
}

func GetConfig() (*Config, error) {
	var pgCfg Config
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
