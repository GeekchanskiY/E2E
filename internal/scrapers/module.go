package scrapers

import (
	"finworker/internal/scrapers/myfin"
	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(myfin.New),
		fx.Invoke(myfin.GetCurrencies),
	)
}
