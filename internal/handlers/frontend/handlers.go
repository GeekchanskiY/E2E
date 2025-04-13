package frontend

import (
	"go.uber.org/fx"

	"finworker/internal/handlers/frontend/base"
	"finworker/internal/handlers/frontend/finance"
	"finworker/internal/handlers/frontend/permissions"
	"finworker/internal/handlers/frontend/work"
)

func Construct() fx.Option {
	return fx.Provide(
		base.New,
		finance.New,
		work.New,
		permissions.New,
	)
}
