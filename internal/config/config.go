package config

import (
	"context"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/migmatore/emivn-tg-bot/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"8081"`
	}
	DBConnection struct {
		Username string `env:"DB_USERNAME" env-default:"migmatore"`
		Password string `env:"DB_PASSWORD" env-default:"root"`
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		DB       string `env:"DB_NAME" env-default:"testdb"`
		MaxConns string `env:"DB_MAX_CONNS" env-default:"50"`
	}
	AppConfig struct {
		LogLevel string `env:"LOG_LEVEL" env-default:"info"`
	}
}

var instance *Config
var once sync.Once

func GetConfig(ctx context.Context) *Config {
	once.Do(func() {
		logging.GetLogger(ctx).Info("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			var title = "Bakery management system"
			help, _ := cleanenv.GetDescription(instance, &title)

			logging.GetLogger(ctx).Info(help)
			logging.GetLogger(ctx).Fatal(err)
		}
	})

	return instance
}
