package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/ios116/frame/config"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	logger, _ := zap.NewDevelopment()
	logger.Debug("This is a DEBUG message")
	logger.Info("This is an INFO message")
	//server := httpserver.HttpServer{
	//	Port:   cfg.Port,
	//	Host:   cfg.Host,
	//}
	//server.Start()
	//logger.DPanic("This is PANIC")
}
