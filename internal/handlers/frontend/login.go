package frontend

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.login")
	if r.Method == http.MethodGet {

		// maybe I should move this logic to controller
		user, ok := r.Context().Value("user").(string)
		if !ok {
			h.logger.Error("frontend.login: user is not a string")

			http.Error(w, "cant get user session data", http.StatusInternalServerError)
			return
		}

		if user != "undefined" {
			http.Redirect(w, r, "/", http.StatusSeeOther)

			return
		}

		html, err := h.controller.Login(r.Context())
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", nil)
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		html, err := h.controller.LoginForm(r.Context(), username, password)
		if err != nil {
			if html == nil {
				h.logger.Error("frontend.login", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = html.ExecuteTemplate(w, "base", map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		if html == nil {
			h.logger.Error("frontend.login", zap.Error(frontend.ErrTemplateNotGenerated))
			http.Error(w, frontend.ErrTemplateNotGenerated.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", nil)
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
