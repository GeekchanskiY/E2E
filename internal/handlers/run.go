package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(h *Handler) error {
	r := chi.NewRouter()
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return http.ListenAndServe(":8080", r)
}
