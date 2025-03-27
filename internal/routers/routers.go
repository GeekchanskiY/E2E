package routers

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"finworker/internal/handlers"
)

type Router struct {
	logger   *zap.Logger
	handlers handlers.Handlers
	config   Config

	mux *chi.Mux
}

func New(logger *zap.Logger, handlers handlers.Handlers, config Config) *Router {
	mux := chi.NewRouter()
	return &Router{
		logger:   logger,
		handlers: handlers,
		config:   config,
		mux:      mux,
	}
}
