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

	server := httpserver.HTTPServer{
		Port:   cfg.Port,
		Host:   cfg.Host,
		Logger: logger,
	}
	server.Start()
	//logger.DPanic("This is PANIC")
}
