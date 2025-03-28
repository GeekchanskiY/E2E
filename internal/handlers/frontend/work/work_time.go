package work

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"
)

func (h *handler) WorkTime(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_wallet.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.WorkTime(r.Context())
		if err != nil {

			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
	case http.MethodPost:
		var (
			worker, workId int64

			err error
		)

		worker, err = strconv.ParseInt(r.Context().Value("userId").(string), 10, 64)
		if err != nil {
			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		workId, err = strconv.ParseInt(r.PostFormValue("work_id"), 10, 64)
		if err != nil {
			h.logger.Error("frontend.create_wallet.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		html, templateData, err := h.controller.WorkTimeForm(r.Context(), workId, worker)
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

		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}

	return
}
