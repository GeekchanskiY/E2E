package finance

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"finworker/internal/models"
)

func (h *handler) CreateOperation(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_operation.handler", zap.String("method", r.Method))
	walletID, err := strconv.ParseInt(chi.URLParam(r, "walletId"), 10, 64)
	if err != nil {
		h.logger.Error("frontend.create_operation.handler", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreateOperation(r.Context(), walletID)
		if err != nil {

			h.logger.Error("frontend.create_operation.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
		if err != nil {
			h.logger.Error("frontend.create_operation.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	case http.MethodPost:
		amount, err := strconv.ParseFloat(r.PostFormValue("amount"), 64)
		if err != nil {
			h.logger.Error("frontend.create_operation.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		operationGroupID, err := strconv.ParseInt(r.PostFormValue("operation_group"), 10, 64)
		if err != nil {
			h.logger.Error("frontend.create_operation.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		html, templateData, err := h.controller.CreateOperationForm(r.Context(), &models.Operation{
			OperationGroupID: operationGroupID,
			Time:             time.Now(),
			IsMonthly:        false,
			IsConfirmed:      true,
			Amount:           amount,
		}, walletID)
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.create_operation.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.create_operation.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, fmt.Sprintf("/finance/wallet/%d", walletID), http.StatusSeeOther)

		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

		return
	}
}
