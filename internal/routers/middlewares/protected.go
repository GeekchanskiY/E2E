package middlewares

import (
	"net/http"

	"finworker/internal/config"
)

func Protected(loginRedirect bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			user, ok := req.Context().Value(config.UsernameContextKey).(string)
			if !ok || user == "undefined" {
				// for FE better usability
				if loginRedirect {
					http.Redirect(w, req, "/login", http.StatusSeeOther)
					return
				}

				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, req)
		})
	}
}
