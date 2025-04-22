package permissions

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/models"
)

func (h *handler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.create_permission.handler", zap.String("event", "got request"))

	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.CreatePermissionGroup(r.Context())
		if err != nil {
			h.logger.Error("frontend.create_permission.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		if err = html.ExecuteTemplate(w, "base", templateData); err != nil {
			h.logger.Error("frontend.create_permission.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	case http.MethodPost:
		name := r.PostFormValue("name")

		html, templateData, err := h.controller.CreatePermissionGroupForm(r.Context(), &models.PermissionGroup{
			Name: name,
		})
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.create_permission.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.create_permission.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/permissions", http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
