package routers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers"
)

type Router struct {
	logger      *zap.Logger
	controllers *controllers.Controllers
	config      Config
}

func NewHandler(logger *zap.Logger, controllers *controllers.Controllers, config Config) *Router {
	return &Router{
		logger:      logger,
		controllers: controllers,
		config:      config,
	}
}
