package frontend

import (
	"net/http"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("frontend.register.handler", zap.String("event", "got request"))
	if r.Method == http.MethodGet {
		h.logger.Debug(
			"frontend.register.handler",
			zap.String("event", "got request"),
			zap.String("method", "GET"),
		)

		html, templateData, err := h.controllers.Base().Register(r.Context())
		if err != nil {
			h.logger.Error("frontend.register", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = html.ExecuteTemplate(w, "base", templateData)
		if err != nil {
			h.logger.Error("frontend.register", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	} else if r.Method == http.MethodPost {
		h.logger.Debug(
			"frontend.register.handler",
			zap.String("event", "got request"),
			zap.String("method", "POST"),
		)
		username := r.PostFormValue("username")
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")
		repeatPassword := r.PostFormValue("repeat_password")
		gender := r.PostFormValue("gender")
		birthday := r.PostFormValue("birthday")
		bank := r.PostFormValue("bank")
		salary := r.PostFormValue("salary")
		currency := r.PostFormValue("currency")
		payday := r.PostFormValue("payday")

		html, templateData, token, salt, err := h.controllers.Base().RegisterForm(r.Context(), username, name, password, repeatPassword, gender, birthday, bank, salary, currency, payday)
		if err != nil {
			if html == nil {
				h.logger.Error("frontend.register", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = html.ExecuteTemplate(w, "base", templateData)
			if err != nil {
				h.logger.Error("frontend.register", zap.Error(err))
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if token == "" {
			h.logger.Error("frontend.register.handler", zap.Error(frontend.ErrTemplateNotGenerated))

			http.Error(w, frontend.ErrTemplateNotGenerated.Error(), http.StatusInternalServerError)

			return
		}

		h.logger.Info(
			"frontend.register.handler",
			zap.String("event", "user registered"),
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

		return
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
