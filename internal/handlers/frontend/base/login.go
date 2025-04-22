package base

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/config"
	"finworker/internal/controllers/frontend"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.logger.Debug(
			"frontend.login.handler",
			zap.String("event", "got request"),
			zap.String("method", "GET"),
		)

		_, ok := r.Context().Value(config.UsernameContextKey).(string)

		if ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)

			return
		}

		html, templateData, err := h.controller.Login(r.Context())
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	case http.MethodPost:
		h.logger.Debug(
			"frontend.login.handler",
			zap.String("event", "got request"),
			zap.String("method", "POST"),
		)

		err := r.ParseForm()
		if err != nil {
			h.logger.Error("frontend.login", zap.Error(err))

			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		html, templateData, token, salt, err := h.controller.LoginForm(r.Context(), username, password)
		if err != nil {
			if html == nil {
				h.logger.Error("frontend.login", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			err = html.ExecuteTemplate(w, "base", templateData)
			if err != nil {
				h.logger.Error("frontend.login", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)

				return
			}

			return
		}

		if token == "" {
			h.logger.Error("frontend.login", zap.Error(frontend.ErrTemplateNotGenerated))

			http.Error(w, frontend.ErrTemplateNotGenerated.Error(), http.StatusInternalServerError)

			return
		}

		h.logger.Info(
			"frontend.login.handler",
			zap.String("event", "user logged in"),
			zap.String("username", username),
		)

		authCookie := http.Cookie{
			Name:     "user",
			Value:    token,
			Path:     "/",
			MaxAge:   7200, // 120 min
			HttpOnly: false,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		saltCookie := http.Cookie{
			Name:     "salt",
			Value:    salt,
			Path:     "/",
			MaxAge:   7200,
			HttpOnly: false,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, &authCookie)
		http.SetCookie(w, &saltCookie)

		http.Redirect(w, r, "/me", http.StatusSeeOther)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
