package frontend

import (
	"net/http"

	"finworker/internal/handlers/frontend/utils"
	"go.uber.org/zap"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.me")

	html, err := h.controller.User(r.Context(), r.Context().Value("user").(string))
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateData := utils.BuildDefaultDataMapFromContext(r.Context())

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
