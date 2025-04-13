package repositories

import (
	"go.uber.org/fx"
)

func Construct(repositories *Repositories) fx.Option {
	return fx.Options(
		fx.Provide(repositories.GetBanks()),
		fx.Provide(repositories.GetUsers()),
		fx.Provide(repositories.GetCurrencyStates()),
		fx.Provide(repositories.GetDistributors()),
		fx.Provide(repositories.GetOperations()),
		fx.Provide(repositories.GetOperationGroups()),
		fx.Provide(repositories.GetPermissionGroups()),
		fx.Provide(repositories.GetWallets()),
		fx.Provide(repositories.GetUserPermissions()),
		fx.Provide(repositories.GetWork()),
	)
}
