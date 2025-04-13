package routers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"finworker/internal/config"
	"finworker/internal/handlers/frontend/base"
	"finworker/internal/handlers/frontend/finance"
	"finworker/internal/handlers/frontend/permissions"
	"finworker/internal/handlers/frontend/work"
	"finworker/internal/handlers/users"
)

type Router struct {
	logger             *zap.Logger
	baseHandler        base.Handler
	financeHandler     finance.Handler
	permissionsHandler permissions.Handler
	workHandler        work.Handler
	usersHandler       users.Handler

	config *config.Config

	mux *chi.Mux
}

func New(
	logger *zap.Logger,
	config *config.Config,
	baseHandler base.Handler,
	financeHandler finance.Handler,
	permissionsHandler permissions.Handler,
	workHandler work.Handler,
	usersHandler users.Handler,
) *Router {
	mux := chi.NewRouter()
	return &Router{
		logger:             logger,
		baseHandler:        baseHandler,
		financeHandler:     financeHandler,
		permissionsHandler: permissionsHandler,
		workHandler:        workHandler,
		usersHandler:       usersHandler,
		config:             config,
		mux:                mux,
	}
}
