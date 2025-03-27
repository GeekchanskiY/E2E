package frontend

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
	"finworker/internal/handlers/frontend/base"
	"finworker/internal/handlers/frontend/finance"
)

type Handlers interface {
	Base() base.Handler
	Finance() finance.Handler
}

type handlers struct {
	base    base.Handler
	finance finance.Handler
}

func New(logger *zap.Logger, controllers frontend.Controllers) Handlers {
	baseHandler := base.New(logger, controllers.Base())
	financeHandler := finance.New(logger, controllers.Finance())
	return &handlers{
		base:    baseHandler,
		finance: financeHandler,
	}
}

func (h *handlers) Base() base.Handler {
	return h.base
}

func (h *handlers) Finance() finance.Handler {
	return h.finance
}
