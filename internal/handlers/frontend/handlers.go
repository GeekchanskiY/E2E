package frontend

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
	"finworker/internal/handlers/frontend/base"
	"finworker/internal/handlers/frontend/finance"
	"finworker/internal/handlers/frontend/work"
)

type Handlers interface {
	Base() base.Handler
	Finance() finance.Handler
	Work() work.Handler
}

type handlers struct {
	base    base.Handler
	finance finance.Handler
	work    work.Handler
}

func New(logger *zap.Logger, controllers frontend.Controllers) Handlers {
	baseHandler := base.New(logger, controllers.Base())
	financeHandler := finance.New(logger, controllers.Finance())
	workHandler := work.New(logger, controllers.Work())
	return &handlers{
		base:    baseHandler,
		finance: financeHandler,
		work:    workHandler,
	}
}

func (h *handlers) Base() base.Handler {
	return h.base
}

func (h *handlers) Finance() finance.Handler {
	return h.finance
}

func (h *handlers) Work() work.Handler {
	return h.work
}
