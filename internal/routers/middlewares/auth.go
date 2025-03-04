package middlewares

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func Auth(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			tokenStr, err := req.Cookie("user")
			if tokenStr == nil || err != nil {
				next.ServeHTTP(w, req)

				return
			}

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenStr.Value, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, err.Error(), http.StatusForbidden)

				return
			}

			username, okUsername := claims["user"].(map[string]interface{})
			userId, okUserId := claims["id"].(float64)

			if !okUsername || !okUserId {
				http.Error(w, "Invalid token claims", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(req.Context(), "user", username["username"])
			ctx = context.WithValue(ctx, "userId", int64(userId))
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
