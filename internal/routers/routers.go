package routers

import (
	"go.uber.org/zap"

	"finworker/internal/handlers"
)

type Router struct {
	logger   *zap.Logger
	handlers *handlers.Handlers
	config   Config
}

func NewRouter(logger *zap.Logger, handlers *handlers.Handlers, config Config) *Router {
	return &Router{
		logger:   logger,
		handlers: handlers,
		config:   config,
	}
}
