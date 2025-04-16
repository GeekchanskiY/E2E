package permissions

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.list.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.List(r.Context())
		if err != nil {
			h.logger.Error("frontend.list.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err = html.ExecuteTemplate(w, "base", templateData); err != nil {
			h.logger.Error("frontend.list.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
