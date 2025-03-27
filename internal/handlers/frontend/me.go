package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.me.handler", zap.String("event", "got request"))

	html, templateData, err := h.controller.Base().User(r.Context(), r.Context().Value("user").(string))
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.me", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
