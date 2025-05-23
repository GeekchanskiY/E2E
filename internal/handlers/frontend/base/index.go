package base

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.index.handler", zap.String("event", "got request"))

	html, templateData, err := h.controller.Index(r.Context(), r.RemoteAddr)
	if err != nil {
		h.logger.Error("frontend.index", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.index", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
