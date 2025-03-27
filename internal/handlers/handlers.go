package handlers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers"
	"finworker/internal/handlers/frontend"
	"finworker/internal/handlers/users"
)

type Handlers struct {
	userHandler      *users.Handler
	frontendHandlers frontend.Handlers
}

func New(logger *zap.Logger, controller *controllers.Controllers) *Handlers {
	return &Handlers{
		userHandler:      users.NewHandler(logger, controller.GetUsers()),
		frontendHandlers: frontend.New(logger, controller.GetFrontend()),
	}
}

func (h *Handlers) GetUsers() *users.Handler {
	return h.userHandler
}

func (h *Handlers) GetFrontend() frontend.Handlers {
	return h.frontendHandlers
}
