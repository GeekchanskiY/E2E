package middlewares

import (
	"context"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token, err := req.Cookie("token")
		if err != nil {
			ctx := context.WithValue(req.Context(), "user", "undefined")
			next.ServeHTTP(w, req.WithContext(ctx))

			return
		}

		ctx := context.WithValue(req.Context(), "user", token.Value)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
