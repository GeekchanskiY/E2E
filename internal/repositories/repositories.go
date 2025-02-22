package repositories

import (
	"github.com/jmoiron/sqlx"

	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currency_states"
	"finworker/internal/repositories/distributors"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/operaton_groups"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
)

type Repositories struct {
	distributors     *distributors.Repository
	operationGroups  *operaton_groups.Repository
	operations       *operations.Repository
	permissionGroups *permission_groups.Repository
	userPermissions  *user_permissions.Repository
	users            *users.Repository
	wallets          *wallets.Repository
	banks            *banks.Repository
	currencyStates   *currency_states.Repository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		distributors:     distributors.New(db),
		operationGroups:  operaton_groups.New(db),
		operations:       operations.New(db),
		permissionGroups: permission_groups.New(db),
		userPermissions:  user_permissions.New(db),
		users:            users.New(db),
		wallets:          wallets.New(db),
		banks:            banks.New(db),
		currencyStates:   currency_states.New(db),
	}
}

func (r *Repositories) GetDistributors() *distributors.Repository {
	return r.distributors
}

func (r *Repositories) GetOperationGroups() *operaton_groups.Repository {
	return r.operationGroups
}

func (r *Repositories) GetOperations() *operations.Repository {
	return r.operations
}

func (r *Repositories) GetPermissionGroups() *permission_groups.Repository {
	return r.permissionGroups
}

func (r *Repositories) GetUserPermissions() *user_permissions.Repository {
	return r.userPermissions
}

func (r *Repositories) GetUsers() *users.Repository {
	return r.users
}

func (r *Repositories) GetWallets() *wallets.Repository {
	return r.wallets
}

func (r *Repositories) GetBanks() *banks.Repository {
	return r.banks
}

func (r *Repositories) GetCurrencyStates() *currency_states.Repository {
	return r.currencyStates
}
