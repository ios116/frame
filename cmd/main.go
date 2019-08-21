package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/ios116/frame/config"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	//server := httpserver.HttpServer{
	//	Port:   cfg.Port,
	//	Host:   cfg.Host,
	//	Router: nil,
	//}
	//server.Start()
	//logger.DPanic("This is PANIC")
}
