package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(h *Handler) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/users", func(r chi.Router) {

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", h.controller.GetUser)
		})
	})

	return http.ListenAndServe(":8080", r)
}
