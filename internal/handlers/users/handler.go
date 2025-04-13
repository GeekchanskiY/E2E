package users

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/users"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller *users.Controller
}

func New(logger *zap.Logger, controller *users.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
