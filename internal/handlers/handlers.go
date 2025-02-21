package handlers

import (
	"finworker/internal/controllers"
	"go.uber.org/zap"
)

type Handler struct {
	logger      *zap.Logger
	controllers *controllers.Controllers
}

func NewHandler(logger *zap.Logger, controllers *controllers.Controllers) *Handler {
	return &Handler{
		logger:      logger,
		controllers: controllers,
	}
}
