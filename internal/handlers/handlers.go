package handlers

import (
	"finworker/internal/controllers"
)

type Handler struct {
	controller controllers.Controller
}

func NewHandler(controller controllers.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}
