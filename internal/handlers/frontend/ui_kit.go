package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) UIKit(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.ui_kit.handler", zap.String("event", "got request"))

	html, templateData, err := h.controller.Base().UIKit(r.Context())
	if err != nil {
		h.logger.Error("frontend.ui_kit", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.ui_kit", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
