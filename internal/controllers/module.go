package controllers

import (
	"go.uber.org/fx"

	"finworker/internal/controllers/frontend"
	"finworker/internal/controllers/media"
	"finworker/internal/controllers/users"
)

func Construct() fx.Option {
	return fx.Options(
		fx.Provide(
			users.New,
			media.New,
		),

		frontend.Construct(),
	)
}
