package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"finworker/docs"
	"finworker/internal/routers/middlewares"
)

func (r *Router) addRoutes() {
	r.addBaseRoutes()

	r.addApiRoutes("/api/v1")
	r.addUserRoutes()

	r.mux.Group(func(m chi.Router) {
		m.Use(middlewares.Protected(true))

		addProtectedFinanceRoutes(r, m)
		addProtectedUserRoutes(r, m)
		addProtectedWorkRoutes(r, m)
		addProtectedPermissionRoutes(r, m)
	})
}

func (r *Router) addBaseRoutes() {
	r.mux.NotFound(r.handlers.GetFrontend().Base().PageNotFound)

	r.mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.mux.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		tmpl := docs.GetTemplate()

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.mux.Get("/", r.handlers.GetFrontend().Base().Index)
	r.mux.Get("/ui_kit", r.handlers.GetFrontend().Base().UIKit)
	r.mux.Get("/faq", r.handlers.GetFrontend().Base().FAQ)
}

func (r *Router) addUserRoutes() {
	r.mux.Get("/login", r.handlers.GetFrontend().Base().Login)
	r.mux.Post("/login", r.handlers.GetFrontend().Base().Login)
	r.mux.Get("/register", r.handlers.GetFrontend().Base().Register)
	r.mux.Post("/register", r.handlers.GetFrontend().Base().Register)
	r.mux.Get("/logout", r.handlers.GetFrontend().Base().Logout)
}

func (r *Router) addApiRoutes(apiPrefix string) {
	r.mux.Route(apiPrefix, func(m chi.Router) {
		m.Route("/users", func(m chi.Router) {
			m.Post("/register", r.handlers.GetUsers().Register)

			m.Route("/{userId}", func(m chi.Router) {
				m.Get("/", r.handlers.GetUsers().Get)
			})
		})
	})
}

func addProtectedFinanceRoutes(r *Router, m chi.Router) {
	m.Route("/finance", func(m chi.Router) {
		m.Get("/", r.handlers.GetFrontend().Finance().Finance)

		m.Get("/create_wallet", r.handlers.GetFrontend().Finance().CreateWallet)
		m.Post("/create_wallet", r.handlers.GetFrontend().Finance().CreateWallet)

		m.Get("/create_distributor", r.handlers.GetFrontend().Finance().CreateDistributor)
		m.Post("/create_distributor", r.handlers.GetFrontend().Finance().CreateDistributor)

		m.Get("/create_operation_group", r.handlers.GetFrontend().Finance().CreateOperationGroup)
		m.Post("/create_operation_group", r.handlers.GetFrontend().Finance().CreateOperationGroup)

		m.Get("/create_operation/{walletId}", r.handlers.GetFrontend().Finance().CreateOperation)
		m.Post("/create_operation/{walletId}", r.handlers.GetFrontend().Finance().CreateOperation)

		m.Get("/wallet/{id}", r.handlers.GetFrontend().Finance().Wallet)
	})
}

func addProtectedWorkRoutes(r *Router, m chi.Router) {
	m.Route("/work", func(m chi.Router) {
		m.Get("/", r.handlers.GetFrontend().Work().WorkTime)
		m.Post("/", r.handlers.GetFrontend().Work().WorkTime)

		m.Get("/create", r.handlers.GetFrontend().Work().CreateWork)
		m.Post("/create", r.handlers.GetFrontend().Work().CreateWork)
	})
}

func addProtectedUserRoutes(r *Router, m chi.Router) {
	m.Get("/me", r.handlers.GetFrontend().Base().Me)
}

func addProtectedPermissionRoutes(r *Router, m chi.Router) {
	m.Route("/permissions", func(m chi.Router) {
		m.Get("/", r.handlers.GetFrontend().Permissions().List)
		m.Get("/group/{id}", r.handlers.GetFrontend().Permissions().PermissionGroup)
	})
}
