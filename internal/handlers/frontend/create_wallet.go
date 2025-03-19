package frontend

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/models"
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
		name := r.PostFormValue("name")
		description := r.PostFormValue("description")
		bank := r.PostFormValue("bank")
		currency := r.PostFormValue("currency")
		permissionGroup := r.PostFormValue("permission")

		isSalary := false
		if r.PostFormValue("is_salary") == "on" {
			isSalary = true
		}

		html, templateData, err := h.controller.CreateWalletForm(r.Context(), models.WalletExtended{
			Name:        name,
			Description: description,
			Permission:  permissionGroup,
			Currency:    models.Currency(currency),
			IsSalary:    isSalary,
			BankName:    bank,
		})
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.create_wallet.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.create_wallet.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/finance", http.StatusSeeOther)

		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}

	return
}
