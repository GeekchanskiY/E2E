package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"finworker/docs"
	"finworker/internal/routers/middlewares"
)

func (r *Router) addRoutes() {
	r.addBaseRoutes()
	r.addMediaRoutes()

	r.addAPIRoutes("/api/v1")
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
	r.mux.NotFound(r.baseHandler.PageNotFound)

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

	r.mux.Get("/", r.baseHandler.Index)
	r.mux.Get("/ui_kit", r.baseHandler.UIKit)
	r.mux.Get("/faq", r.baseHandler.FAQ)
}

func (r *Router) addMediaRoutes() {
	r.mux.Post("/media", r.mediaHandler.Upload)
}

func (r *Router) addUserRoutes() {
	r.mux.Get("/login", r.baseHandler.Login)
	r.mux.Post("/login", r.baseHandler.Login)
	r.mux.Get("/register", r.baseHandler.Register)
	r.mux.Post("/register", r.baseHandler.Register)
	r.mux.Get("/logout", r.baseHandler.Logout)
}

func (r *Router) addAPIRoutes(apiPrefix string) {
	r.mux.Route(apiPrefix, func(m chi.Router) {
		m.Route("/users", func(m chi.Router) {
			m.Post("/register", r.usersHandler.Register)

			m.Route("/{userId}", func(m chi.Router) {
				m.Get("/", r.usersHandler.Get)
			})
		})
	})
}

func addProtectedFinanceRoutes(r *Router, m chi.Router) {
	m.Route("/finance", func(m chi.Router) {
		m.Get("/", r.financeHandler.Finance)

		m.Get("/create_wallet", r.financeHandler.CreateWallet)
		m.Post("/create_wallet", r.financeHandler.CreateWallet)

		m.Get("/create_distributor", r.financeHandler.CreateDistributor)
		m.Post("/create_distributor", r.financeHandler.CreateDistributor)

		m.Get("/create_operation_group", r.financeHandler.CreateOperationGroup)
		m.Post("/create_operation_group", r.financeHandler.CreateOperationGroup)

		m.Get("/create_operation/{walletId}", r.financeHandler.CreateOperation)
		m.Post("/create_operation/{walletId}", r.financeHandler.CreateOperation)

		m.Get("/wallet/{id}", r.financeHandler.Wallet)
	})
}

func addProtectedWorkRoutes(r *Router, m chi.Router) {
	m.Route("/work", func(m chi.Router) {
		m.Get("/", r.workHandler.WorkTime)
		m.Post("/", r.workHandler.WorkTime)

		m.Get("/create", r.workHandler.CreateWork)
		m.Post("/create", r.workHandler.CreateWork)
	})
}

func addProtectedUserRoutes(r *Router, m chi.Router) {
	m.Get("/me", r.baseHandler.Me)
	m.Post("/avatar", r.baseHandler.UploadAvatar)
}

func addProtectedPermissionRoutes(r *Router, m chi.Router) {
	m.Route("/permissions", func(m chi.Router) {
		m.Get("/", r.permissionsHandler.List)

		m.Get("/create", r.permissionsHandler.CreatePermission)
		m.Post("/create", r.permissionsHandler.CreatePermission)

		m.Get("/group/{id}", r.permissionsHandler.PermissionGroup)

		m.Get("/group/{id}/add", r.permissionsHandler.AddUser)
		m.Post("/group/{id}/add", r.permissionsHandler.AddUser)
	})
}
