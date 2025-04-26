package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	Logger *zap.Logger `config:"-"`

	Secret     string `config:"APP_SECRET"`
	Host       string `config:"APP_HOST"`
	Port       int    `config:"APP_PORT"`
	DbName     string `config:"DB_NAME"`
	DbUser     string `config:"DB_USER"`
	DbPassword string `config:"DB_PASSWORD"`
	DbHost     string `config:"DB_HOST"`
	DbPort     string `config:"DB_PORT"`
}

func NewConfig() *Config {
	var (
		logger *zap.Logger
		loader *confita.Loader

		err error
	)

	if logger, err = zap.NewDevelopment(); err != nil {
		panic(err)
	}

	cfg := new(Config)
	cfg.Logger = logger

	if err = godotenv.Load(); err != nil {
		logger.Fatal("error loading .env file", zap.Error(err))
	}

	loader = confita.NewLoader()

	if err = loader.Load(context.Background(), cfg); err != nil {
		logger.Fatal("error loading config", zap.Error(err))
	}

	return cfg
}

func GetLogger(cfg *Config) *zap.Logger {
	return cfg.Logger
}
