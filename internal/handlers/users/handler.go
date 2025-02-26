package users

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/users"
)

type Handler struct {
	logger     *zap.Logger
	controller *users.Controller
}

func NewHandler(logger *zap.Logger, controller *users.Controller) *Handler {
	return &Handler{
		logger:     logger,
		controller: controller,
	}
}
