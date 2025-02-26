package app

import (
	"finworker/internal/controllers"
	"finworker/internal/handlers"
	"finworker/internal/repositories"
	"finworker/internal/routers"
	"finworker/internal/scrapers"
	"finworker/internal/storage"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(

		fx.Provide(
			NewConfig,
			GetDb,     // gets database config from main config instance
			GetLogger, // gets logger from main config instance
			GetRouter,
		),

		// main logic modules & http server
		storage.NewModule(),
		repositories.NewModule(),
		controllers.NewModule(),
		handlers.NewModule(),
		routers.NewModule(),

		// scrapers & periodic tasks
		scrapers.NewModule(),
	)
}
