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
	r.Use(middlewares.Auth(h.config.Secret))

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
	r.Get("/ui_kit", h.handlers.GetFrontend().UIKit)

	// User routes
	r.Get("/login", h.handlers.GetFrontend().Login)
	r.Post("/login", h.handlers.GetFrontend().Login)
	r.Get("/register", h.handlers.GetFrontend().Register)
	r.Get("/logout", h.handlers.GetFrontend().Logout)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.Protected(true))
		r.Get("/finance", h.handlers.GetFrontend().Finance)
		r.Get("/finance/create_wallet", h.handlers.GetFrontend().CreateWallet)
		r.Post("/finance/create_wallet", h.handlers.GetFrontend().CreateWallet)
		r.Get("/finance/wallet/{id}", h.handlers.GetFrontend().Wallet)
		r.Get("/me", h.handlers.GetFrontend().Me)
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", h.handlers.GetUsers().Register)

			r.Route("/{userId}", func(r chi.Router) {
				r.Get("/", h.handlers.GetUsers().Get)
			})
		})
	})

	go func() {
		h.logger.Info(
			"running server",
			zap.String("addr", fmt.Sprintf("%s:%d", h.config.Host, h.config.Port)),
		)

		err := http.ListenAndServe(fmt.Sprintf("%s:%d", h.config.Host, h.config.Port), r)
		if err != nil {
			h.logger.Fatal("failed to start http server", zap.Error(err))
		}
	}()

	return nil
}
