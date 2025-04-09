package permissions

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/permissions"
)

type Handler interface {
	CreatePermission(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller permissions.Controller
}

func New(logger *zap.Logger, controller permissions.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
