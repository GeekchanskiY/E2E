package handlers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers"
	"finworker/internal/handlers/frontend"
	"finworker/internal/handlers/users"
)

type Handlers interface {
	GetUsers() *users.Handler
	GetFrontend() frontend.Handlers
}

type handlers struct {
	userHandler      *users.Handler
	frontendHandlers frontend.Handlers
}

func New(logger *zap.Logger, controller *controllers.Controllers) Handlers {
	return &handlers{
		userHandler:      users.NewHandler(logger, controller.GetUsers()),
		frontendHandlers: frontend.New(logger, controller.GetFrontend()),
	}
}

func (h *handlers) GetUsers() *users.Handler {
	return h.userHandler
}

func (h *handlers) GetFrontend() frontend.Handlers {
	return h.frontendHandlers
}
