package handlers

import (
	"net/http"

	"finworker/docs"
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
	r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		tmpl := docs.GetTemplate()

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", h.controller.RegisterUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", h.controller.GetUser)
		})
	})

	return http.ListenAndServe(":8080", r)
}
