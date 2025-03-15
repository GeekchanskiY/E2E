package frontend

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) PageNotFound(w http.ResponseWriter, r *http.Request) {
	h.logger.Error("frontend.page_not_found", zap.String("route", r.RequestURI))

	html, templateData, err := h.controller.PageNotFound(r.Context())
	if err != nil {
		h.logger.Error("frontend.page_not_found", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", templateData)
	if err != nil {
		h.logger.Error("frontend.page_not_found", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
