package handlers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers"
	"finworker/internal/handlers/frontend"
	"finworker/internal/handlers/users"
)

type Handlers struct {
	userHandler     *users.Handler
	frontendHandler *frontend.Handler
}

func New(logger *zap.Logger, controller *controllers.Controllers) *Handlers {
	return &Handlers{
		userHandler:     users.NewHandler(logger, controller.GetUsers()),
		frontendHandler: frontend.NewHandler(logger, controller.GetFrontend()),
	}
}

func (h *Handlers) GetUsers() *users.Handler {
	return h.userHandler
}

func (h *Handlers) GetFrontend() *frontend.Handler {
	return h.frontendHandler
}
