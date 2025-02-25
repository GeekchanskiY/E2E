package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"finworker/internal/static"

	"finworker/docs"
)

func Run(h *Router) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

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

	r.Get("/", h.controllers.GetFrontend().Index)
	r.Get("/finance", h.controllers.GetFrontend().Finance)

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", h.controllers.GetUsers().RegisterUser)

			r.Route("/{userId}", func(r chi.Router) {
				r.Get("/", h.controllers.GetUsers().GetUser)
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
