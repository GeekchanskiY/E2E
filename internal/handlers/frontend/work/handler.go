package work

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/work"
)

type Handler interface {
	CreateWork(w http.ResponseWriter, r *http.Request)
	WorkTime(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller work.Controller
}

func New(logger *zap.Logger, controller work.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
