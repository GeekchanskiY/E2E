package frontend

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"
	"finworker/internal/handlers/frontend/base"
	"finworker/internal/handlers/frontend/finance"
	"finworker/internal/handlers/frontend/permissions"
	"finworker/internal/handlers/frontend/work"
)

type Handlers interface {
	Base() base.Handler
	Finance() finance.Handler
	Work() work.Handler
	Permissions() permissions.Handler
}

type handlers struct {
	base        base.Handler
	finance     finance.Handler
	work        work.Handler
	permissions permissions.Handler
}

func New(logger *zap.Logger, controllers frontend.Controllers) Handlers {
	baseHandler := base.New(logger, controllers.Base())
	financeHandler := finance.New(logger, controllers.Finance())
	workHandler := work.New(logger, controllers.Work())
	permissionHandler := permissions.New(logger, controllers.Permissions())
	return &handlers{
		base:        baseHandler,
		finance:     financeHandler,
		work:        workHandler,
		permissions: permissionHandler,
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

func (h *handlers) Permissions() permissions.Handler {
	return h.permissions
}
