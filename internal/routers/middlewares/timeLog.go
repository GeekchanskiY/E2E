package middlewares

import (
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

var excludedRoutes = []string{"/static"}

func Log(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			for _, route := range excludedRoutes {
				if strings.HasPrefix(req.URL.Path, route) {
					next.ServeHTTP(w, req)

					return
				}
			}

			start := time.Now()
			next.ServeHTTP(w, req)
			log.Debug("request executed", zap.String("method", req.Method), zap.String("url", req.URL.String()), zap.Duration("duration", time.Since(start)))
		})
	}
}
