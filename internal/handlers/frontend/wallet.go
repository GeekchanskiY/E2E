package frontend

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (h *Handler) Wallet(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.register.handler", zap.String("event", "got request"))

	walletId := chi.URLParam(r, "id")
	if walletId == "" {
		h.logger.Error("frontend.wallet.register.handler: walletId is empty")
		http.Error(w, "walletId is empty", http.StatusBadRequest)
		return
	}

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
