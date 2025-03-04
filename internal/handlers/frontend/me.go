package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.me")

	html, err := h.controller.User(r.Context(), "geekchanskiy")
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", nil)
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
