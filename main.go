package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/ios116/frame/config"
	"go.uber.org/zap"
	"time"
)

func main() {

	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	logger,_ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Info("This is INFO")
	logger.Debug("This is DEBUG")
	logger.Error("This is ERROR")

	logger,_ = zap.NewDevelopment()
	defer logger.Sync()
	slogger := logger.Sugar()
	slogger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	slogger.Info("This is INFO")
	slogger.Debug("This is DEBUG")
	slogger.Error("This is ERROR")





	//logger.DPanic("This is PANIC")
}
