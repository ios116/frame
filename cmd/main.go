package main

import (
	"github.com/ios116/frame/config"
	"github.com/ios116/frame/httpserver"
	"go.uber.org/zap"
)

func createLogger(build string) (logger *zap.Logger, err error) {
	if build == "dev" {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	return logger, err
}

func main() {

	cfg := &config.Config{}
	cfg.CreateConfig()
	logger, _ := cfg.CreateLogger()
	defer logger.Sync()

	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	logger.Error("This is an ERROR message")
	logger.Warn("This is a WARN messages")

	server := httpserver.HttpServer{
		Port:   cfg.Port,
		Host:   cfg.Host,
	}
	server.Start()
	//logger.DPanic("This is PANIC")
}
