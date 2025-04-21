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
func GetLogger(cfg *Config) *zap.Logger {
	return cfg.Logger
}
