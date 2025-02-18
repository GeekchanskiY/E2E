package app

import (
	"finworker/internal/controllers"
	"finworker/internal/handlers"
	"finworker/internal/repository"
	"finworker/internal/storage"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(

		fx.Provide(
			NewConfig,
			GetDb,     // gets database config from main config instance
			GetLogger, // gets logger from main config instance
		),

		storage.NewModule(),
		repository.NewModule(),
		controllers.NewModule(),
		handlers.NewModule(),
	)
}
