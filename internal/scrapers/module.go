package scrapers

import (
	"go.uber.org/fx"

	"finworker/internal/scrapers/myfin"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(myfin.New),
		fx.Invoke(myfin.RunPeriodicInspection),
	)
}
