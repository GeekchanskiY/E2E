package routers

import (
	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(New),
		fx.Invoke(Run),
	)
}
