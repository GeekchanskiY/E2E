package handlers

import (
	"finworker/internal/controllers"
)

type Handler struct {
	controllers *controllers.Controllers
}

func NewHandler(controllers *controllers.Controllers) *Handler {
	return &Handler{
		controllers: controllers,
	}
}
