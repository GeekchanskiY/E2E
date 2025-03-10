package frontend

import (
	"net/http"

	"finworker/internal/models"
	"go.uber.org/zap"
)

func (h *Handler) CreateWallet(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_wallet.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreateWallet(r.Context())
		if err != nil {

			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
	case http.MethodPost:
		html, templateData, err := h.controller.CreateWalletForm(r.Context(), models.WalletExtended{})
		if err != nil {
			h.logger.Error("frontend.create_wallet.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
		if err != nil {
			h.logger.Error("frontend.create_wallet.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}

	return
}
