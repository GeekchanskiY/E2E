package finance

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (h *handler) Wallet(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.wallet.handler", zap.String("event", "got request"))

	walletId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if walletId == 0 || err != nil {
		h.logger.Error("frontend.wallet.handler: walletId is empty", zap.Error(err))
		http.Error(w, "walletId is empty", http.StatusBadRequest)
		return
	}

	html, templateData, err := h.controller.Wallet(r.Context(), walletId)
	if err != nil {
		h.logger.Error("frontend.wallet", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.wallet", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
