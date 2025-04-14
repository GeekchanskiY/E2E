package app

import (
	"go.uber.org/fx"

	"finworker/internal/config"
	"finworker/internal/controllers"
	"finworker/internal/handlers"
	"finworker/internal/repositories"
	"finworker/internal/routers"
	"finworker/internal/scrapers"
	"finworker/internal/storage"
)

func NewApp() *fx.App {
	return fx.New(

		fx.Provide(
			config.NewConfig,
			config.GetLogger,
		),

		storage.NewModule(),
		repositories.Construct(),
		controllers.Construct(),
		handlers.NewModule(),

		routers.NewModule(),

		scrapers.NewModule(),
	)
}
