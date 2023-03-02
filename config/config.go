package config

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port     string `env:"SRV_PORT" envDefault:"8000"`
	Host     string `env:"SRV_HOST" envDefault:"localhost"`
	PGUser   string `env:"PG_USER" envDefault:"postgres"`
	PGPass   string `env:"PG_PASSWORD" envDefault:"example"`
	PGPort   string `env:"PG_PORT" envDefault:"5432"`
	PGHost   string `env:"PG_HOST" envDefault:"localhost"`
	Timeout  int    `env:"TIMEOUT" envDefault:"5"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`
}

var (
	config Config
	once   sync.Once
)

// NewConfig create service config
func NewConfig() *Config {
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			log.Fatalf("Can't load configuration: %s", err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Load config successful %v", string(configBytes))
	})
	return &config
}
