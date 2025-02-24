package routers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers"
)

type Router struct {
	logger      *zap.Logger
	controllers *controllers.Controllers
}

func NewHandler(logger *zap.Logger, controllers *controllers.Controllers) *Router {
	return &Router{
		logger:      logger,
		controllers: controllers,
	}
}
