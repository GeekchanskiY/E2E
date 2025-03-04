package frontend

import (
	"net/http"

	"finworker/internal/handlers/frontend/utils"
	"go.uber.org/zap"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.index.handler", zap.String("event", "got request"))

	html, err := h.controller.Index(r.Context())
	if err != nil {
		h.logger.Error("frontend.index", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateData := utils.BuildDefaultDataMapFromContext(r.Context())

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.index", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
