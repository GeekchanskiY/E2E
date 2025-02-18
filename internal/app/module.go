package app

import (
	"finworker/internal/repository"
	"go.uber.org/fx"

	"finworker/internal/storage"
)

func NewApp() *fx.App {
	return fx.New(

		fx.Provide(
			NewConfig,
			GetDb,
			GetLogger,
		),
		repository.NewModule(),

		storage.NewModule(),
	)
}
