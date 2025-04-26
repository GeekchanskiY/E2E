package repositories

import (
	"go.uber.org/fx"

	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currencyStates"
	"finworker/internal/repositories/distributors"
	"finworker/internal/repositories/operationGroups"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/registry"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
	"finworker/internal/repositories/works"
)

func Construct() fx.Option {
	return fx.Provide(
		distributors.New,
		operationGroups.New,
		operations.New,
		user_permissions.New,
		permission_groups.New,
		users.New,
		wallets.New,
		banks.New,
		currencyStates.New,
		registry.New,
		works.New,
	)
}
