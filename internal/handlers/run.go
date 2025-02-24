package handlers

import (
	"io/fs"
	"net/http"

	"finworker/internal/static"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"finworker/docs"
)

func Run(h *Handler) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	// static handling
	staticFilesFs, _ := fs.Sub(static.Fs, "files")
	fileserver := http.FileServer(http.FS(staticFilesFs))

	r.Handle("/static/*", http.StripPrefix("/static", fileserver))

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
		r.Post("/register", h.controllers.GetUsers().RegisterUser)

		r.Route("/{userId}", func(r chi.Router) {
			r.Get("/", h.controllers.GetUsers().GetUser)
		})
	})

	// used to avoid fx lock
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			h.logger.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	return nil
}
