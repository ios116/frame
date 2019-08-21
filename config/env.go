package config

// Config  for app
type Config struct {
	Port int    `env:"APP_PORT" envDefault:"3000"`
	Host string `env:"APP_HOST,required"`
	Db string `env:"APP_DB"`
}
