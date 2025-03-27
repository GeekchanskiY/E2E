package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) FAQ(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.FAQ.handler", zap.String("event", "got request"))

	html, templateData, err := h.controllers.Base().FAQ(r.Context())
	if err != nil {
		h.logger.Error("frontend.FAQ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.FAQ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
