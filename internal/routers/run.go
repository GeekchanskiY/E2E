package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"finworker/docs"
	"finworker/internal/routers/middlewares"
	"finworker/internal/static"
)

func Run(h *Router) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middlewares.Auth)

	// static handling
	fileServer := http.FileServer(http.FS(static.Fs))

	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

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

	r.Get("/", h.handlers.GetFrontend().Index)
	r.Get("/finance", h.handlers.GetFrontend().Finance)
	r.Get("/login", h.handlers.GetFrontend().Login)
	r.Post("/login", h.handlers.GetFrontend().Login)
	r.Get("/register", h.handlers.GetFrontend().Register)

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", h.handlers.GetUsers().Register)

			r.Route("/{userId}", func(r chi.Router) {
				r.Get("/", h.handlers.GetUsers().Get)
			})
		})
	})

	// used to avoid fx lock
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", h.config.Host, h.config.Port), r)
		if err != nil {
			h.logger.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	return nil
}
