package app

import (
	"go.uber.org/fx"

	"finworker/internal/controllers"
	"finworker/internal/handlers"
	"finworker/internal/repositories"
	"finworker/internal/routers"
	"finworker/internal/scrapers"
	"finworker/internal/storage"
)

func NewApp() *fx.App {
	return fx.New(
		storage.NewModule(),
		controllers.NewModule(),
		handlers.NewModule(),
		routers.NewModule(),

		scrapers.NewModule(),

		fx.Provide(
			NewConfig,
			GetDbConfig,
			GetLogger,
			GetControllersConfig,
			GetRouterConfig,

			repositories.NewRepositories,
			repositories.Construct,
		),
	)
}
