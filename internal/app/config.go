package app

import (
	"context"

	"github.com/heetch/confita"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"finworker/internal/routers"
	"finworker/internal/storage"
)

type Config struct {
	Logger *zap.Logger `config:"-"`
	Db     storage.Config
	Router routers.Config
}

func NewConfig() *Config {
	logger, _ := zap.NewDevelopment()
	cfg := Config{
		Logger: logger,
	}

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	loader := confita.NewLoader()
	err = loader.Load(context.Background(), &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func GetDb(cfg *Config) storage.Config {
	return cfg.Db
}
func GetRouter(cfg *Config) routers.Config { return cfg.Router }
func GetLogger(cfg *Config) *zap.Logger {
	return cfg.Logger
}
