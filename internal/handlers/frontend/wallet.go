package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) Wallet(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.register.handler", zap.String("event", "got request"))

	html, templateData, err := h.controller.Wallet(r.Context(), 1)
	if err != nil {
		h.logger.Error("frontend.register", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.register", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
