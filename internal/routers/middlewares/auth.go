package middlewares

import (
	"context"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "user", "user")
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
