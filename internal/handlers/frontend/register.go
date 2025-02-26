package frontend

import (
	"net/http"

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

	err = html.ExecuteTemplate(w, "base", nil)
	if err != nil {
		h.logger.Error("frontend.register", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
