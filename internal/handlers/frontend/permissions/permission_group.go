package permissions

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (h *handler) PermissionGroup(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.permission_group.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		permissionGroupId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if permissionGroupId == 0 || err != nil {
			h.logger.Error("frontend.permission_group.handler: permission_group_id is empty", zap.Error(err))
			http.Error(w, "permission_group is empty", http.StatusBadRequest)
			return
		}

		html, templateData, err := h.controller.PermissionGroup(r.Context(), permissionGroupId)
		if err != nil {

			h.logger.Error("frontend.permission_group.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	return
}
