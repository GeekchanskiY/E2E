package base

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/config"
)

func (h *handler) Me(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.me.handler", zap.String("method", r.Method))

	html, templateData, err := h.controller.User(r.Context(), r.Context().Value(config.UsernameContextKey).(string))
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
