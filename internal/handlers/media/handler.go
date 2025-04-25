package media

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/media"
)

type Handler interface {
	Upload(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger     *zap.Logger
	controller media.Controller
}

func New(logger *zap.Logger, controller media.Controller) Handler {
	return &handler{
		logger:     logger,
		controller: controller,
	}
}
