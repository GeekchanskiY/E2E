package frontend

import (
	"net/http"

	"finworker/internal/handlers/frontend/utils"
	"go.uber.org/zap"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.register")
	html, err := h.controller.Register(r.Context())
	if err != nil {
		h.logger.Error("frontend.register", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateData := utils.BuildDefaultDataMapFromContext(r.Context())

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.register", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
