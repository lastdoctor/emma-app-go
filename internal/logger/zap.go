package logger

import (
	"flag"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	environment := flag.String("environment", "development", "set logger environment")
	flag.Parse()

	if *environment == "production" {
		SetDevelopmentLogger()
	} else {
		SetProductionLogger()
	}
}

func Logger() *zap.Logger {
	return logger
}

func SetProductionLogger() {
	logger, _ = zap.NewProduction()
}

func SetDevelopmentLogger() {
	logger, _ = zap.NewDevelopment()
}
