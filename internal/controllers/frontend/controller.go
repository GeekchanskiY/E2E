package frontend

import (
	"go.uber.org/fx"

	"finworker/internal/controllers/frontend/base"
	"finworker/internal/controllers/frontend/finance"
	"finworker/internal/controllers/frontend/permissions"
	"finworker/internal/controllers/frontend/work"
)

func Construct() fx.Option {
	return fx.Provide(
		base.New,
		finance.New,
		permissions.New,
		work.New,
	)
}
