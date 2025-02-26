package frontend

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
)

type Handler struct {
	logger     *zap.Logger
	controller *frontend.Controller
}

func NewHandler(logger *zap.Logger, controller *frontend.Controller) *Handler {
	return &Handler{
		logger:     logger,
		controller: controller,
	}
}
