package config

type Config struct r
	Port int    `env:"APP_PORT" envDefault:"3000"`
	Host string `env:"APP_HOST,required"`
}
