package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

type (
	Config struct {
		Environment string `mapstructure:"Environment"`
		HTTP        HTTPConfig
		Limiter     LimiterConfig
	}
	HTTPConfig struct {
		Host               string        `mapstructure:"HOST"`
		Port               string        `mapstructure:"PORT"`
		ReadTimeout        time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout       time.Duration `mapstructure:"WRITE_TIMEOUT"`
		MaxHeaderMegabytes int           `mapstructure:"MAX_HEADER_BYTES"`
	}

	LimiterConfig struct {
		RPS   int           `mapstructure:"RPS"`
		Burst int           `mapstructure:"BURST"`
		TTL   time.Duration `mapstructure:"TTL"`
	}
)

func ConfigInit(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("development")
	viper.SetConfigName("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		zap.Error(err)
		return nil, err
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	return &cfg, nil
}
