package middlewares

import "net/http"

func Protected(loginRedirect bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			user := req.Context().Value("user")
			if user == nil || user == "undefined" {
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
