package frontend

import (
	"net/http"

	"finworker/internal/handlers/frontend/utils"
	"go.uber.org/zap"
)

func (h *Handler) Finance(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.finance")
	html, err := h.controller.Finance(r.Context())
	if err != nil {
		h.logger.Error("frontend.finance", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templateData := utils.BuildDefaultDataMapFromContext(r.Context())

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.finance", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
