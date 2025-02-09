package app

import (
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

		storage.NewModule(),
	)
}
