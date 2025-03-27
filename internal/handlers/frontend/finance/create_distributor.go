package finance

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (h *handler) CreateDistributor(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_distributor.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreateDistributor(r.Context())
		if err != nil {

			h.logger.Error("frontend.create_distributor.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
	case http.MethodPost:
		name := r.PostFormValue("name")

		sourceWallet, err := strconv.ParseInt(r.PostFormValue("source_wallet"), 10, 64)
		if err != nil {
			sourceWallet = 0
		}
		targetWallet, err := strconv.ParseInt(r.PostFormValue("target_wallet"), 10, 64)
		if err != nil {
			targetWallet = 0
		}

		percent, err := strconv.ParseFloat(r.PostFormValue("percent"), 64)
		if err != nil {
			percent = 0.0
		}

		html, templateData, err := h.controller.CreateDistributorForm(r.Context(), models.Distributor{
			Name:           name,
			SourceWalletId: sourceWallet,
			TargetWalletId: targetWallet,
			Percent:        percent,
		})
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.create_distributor.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.create_distributor.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/finance", http.StatusSeeOther)

		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}
}
