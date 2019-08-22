package config

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"log"
)

// Config for app
type Config struct {
	Port  int    `env:"APP_PORT" envDefault:"3000"`
	Host  string `env:"APP_HOST,required"`
	Db    string `env:"APP_DB"`
	Build string `env:"APP_BUILD" envDefault:"dev"`
}

// CreateConfig parsing environment variables
func (c *Config) CreateConfig() {
	if err := env.Parse(c); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

// CreateLogger creating the logger
func (c *Config) CreateLogger() (logger *zap.Logger, err error) {

	switch c.Build {
	case "dev":
		logger, err = zap.NewDevelopment()
	default:
		logger, err = zap.NewProduction()
	}
	return logger, err
}
