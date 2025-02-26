package frontend

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("frontend.login")
	if r.Method == http.MethodGet {
		user, ok := r.Context().Value("user").(string)
		if ok {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		fmt.Println(user)

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
		html, err := h.controller.LoginForm(r.Context())
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
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
