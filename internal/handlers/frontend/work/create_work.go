package work

import (
	"database/sql"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"finworker/internal/config"
	"finworker/internal/models"
)

func (h *handler) CreateWork(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_wallet.handler", zap.String("event", "got request"))

	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreateWork(r.Context())
		if err != nil {
			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err = html.ExecuteTemplate(w, "base", templateData); err != nil {
			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	case http.MethodPost:
		var (
			name             string
			hourlyRateSource float64
			hourlyRate       sql.NullFloat64

			err error
		)

		name = r.PostFormValue("name")

		worker, ok := r.Context().Value(config.UserIDContextKey).(int64)
		if !ok {
			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))
			http.Error(w, "cant get userid", http.StatusBadRequest)

			return
		}

		hourlyRateSource, err = strconv.ParseFloat(r.PostFormValue("hourlyRate"), 64)
		hourlyRate = sql.NullFloat64{
			Float64: hourlyRateSource,
			Valid:   true,
		}

		if err != nil {
			hourlyRate = sql.NullFloat64{
				Valid: false,
			}
		}

		html, templateData, err := h.controller.CreateWorkForm(r.Context(), &models.UserWork{
			Name:       name,
			HourlyRate: hourlyRate,
			Worker:     worker,
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

			h.logger.Error("frontend.create_work.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/work", http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
