package app

import (
	"context"

	"finworker/internal/controllers"
	"github.com/heetch/confita"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"finworker/internal/routers"
	"finworker/internal/storage"
)

type Config struct {
	Logger      *zap.Logger `config:"-"`
	Controllers controllers.Config
	Db          storage.Config
	Router      routers.Config
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

func GetDbConfig(cfg *Config) storage.Config {
	return cfg.Db
}
func GetRouterConfig(cfg *Config) routers.Config { return cfg.Router }
func GetLogger(cfg *Config) *zap.Logger {
	return cfg.Logger
}
func GetControllersConfig(cfg *Config) controllers.Config { return cfg.Controllers }
