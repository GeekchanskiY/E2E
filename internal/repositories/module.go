package repositories

import (
	"go.uber.org/fx"

	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currency_states"
	"finworker/internal/repositories/distributors"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/operaton_groups"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
	"finworker/internal/repositories/works"
)

func Construct() fx.Option {
	return fx.Provide(
		distributors.New,
		operaton_groups.New,
		operations.New,
		user_permissions.New,
		permission_groups.New,
		users.New,
		wallets.New,
		banks.New,
		currency_states.New,
		works.New,
	)
}
