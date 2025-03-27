package routers

import (
	"github.com/go-chi/chi/v5/middleware"

	"finworker/internal/routers/middlewares"
)

// Setup applies base routes and middlewares
func (r *Router) setup() {
	r.mux.Use(middleware.Recoverer)
	r.mux.Use(middleware.RealIP)
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middlewares.Auth(r.config.Secret))
}
