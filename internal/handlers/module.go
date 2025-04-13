package handlers

import (
	"go.uber.org/fx"

	"finworker/internal/handlers/frontend"
	"finworker/internal/handlers/users"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(users.New),

		frontend.Construct(),
	)
}
