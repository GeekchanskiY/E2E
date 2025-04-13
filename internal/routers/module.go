package routers

import (
	"go.uber.org/fx"
)

const (
	moduleName = "routers"
)

func NewModule() fx.Option {
	return fx.Module(
		moduleName,

		fx.Provide(New),
		fx.Invoke(Run),
	)
}
