package finance

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) Finance(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.finance.handler", zap.String("event", "got request"))
	html, templateData, err := h.controller.Finance(r.Context())
	if err != nil {
		h.logger.Error("frontend.finance", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.finance", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
