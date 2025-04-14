package scrapers

import (
	"go.uber.org/fx"

	"finworker/internal/scrapers/myfin"
)

const (
	moduleName = "scrapers"
)

func NewModule() fx.Option {
	return fx.Module(
		moduleName,

		fx.Provide(myfin.New),

		fx.Invoke(myfin.RunPeriodicScraping),
	)
}
