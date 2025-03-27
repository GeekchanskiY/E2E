package frontend

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
)

type Handler struct {
	logger      *zap.Logger
	controllers frontend.Controllers
}

func NewHandler(logger *zap.Logger, controllers frontend.Controllers) *Handler {
	return &Handler{
		logger:      logger,
		controllers: controllers,
	}
}
