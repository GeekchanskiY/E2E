package finance

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (h *handler) CreateOperationGroup(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_operation_group.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreateOperationGroup(r.Context())
		if err != nil {

			h.logger.Error("frontend.create_operation_group.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
	case http.MethodPost:
		name := r.PostFormValue("name")

		walletId, err := strconv.ParseInt(r.PostFormValue("wallet"), 10, 64)
		if err != nil {
			walletId = 0
		}

		html, templateData, err := h.controller.CreateOperationGroupForm(r.Context(), &models.OperationGroup{
			Name:     name,
			WalletId: walletId,
		})
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.create_operation_group.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.create_operation_group.controller", zap.Error(err))
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
