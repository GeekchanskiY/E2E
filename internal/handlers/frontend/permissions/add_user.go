package permissions

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.add_user.handler", zap.String("event", "got request"))
	switch r.Method {
	case http.MethodGet:
		html, templateData, err := h.controller.AddUser(r.Context())
		if err != nil {

			h.logger.Error("frontend.add_user.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		permissionGroupID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			h.logger.Error("frontend.add_user.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		templateData["permissionGroupId"] = permissionGroupID

		if err = html.ExecuteTemplate(w, "base", templateData); err != nil {
			h.logger.Error("frontend.add_user.handler", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	case http.MethodPost:
		username := r.PostFormValue("username")
		level := r.PostFormValue("level")
		permissionGroupID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			h.logger.Error("frontend.add_user.handler", zap.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		html, templateData, err := h.controller.AddUserForm(r.Context(), username, level, permissionGroupID)
		if err != nil {
			if html != nil {
				err = html.ExecuteTemplate(w, "base", templateData)
				if err != nil {
					h.logger.Error("frontend.add_user.controller", zap.Error(err))
					http.Error(w, err.Error(), http.StatusInternalServerError)

					return
				}

				return
			}

			h.logger.Error("frontend.add_user.controller", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.Redirect(w, r, "/permissions", http.StatusSeeOther)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)

	}
}
