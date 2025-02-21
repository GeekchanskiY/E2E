package repositories

import (
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currency_states"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Distributors     *DistributorsRepository // TODO: make private and create getters and setters
	OperationGroups  *OperationGroupRepository
	Operations       *OperationRepository
	PermissionGroups *PermissionGroupRepository
	UserPermissions  *UserPermissionRepository
	Users            *UserRepository
	Wallets          *WalletRepository
	Banks            *banks.Repository
	CurrencyStates   *currency_states.Repository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Distributors:     NewDistributorsRepository(db),
		OperationGroups:  NewOperationGroupRepository(db),
		Operations:       NewOperationRepository(db),
		PermissionGroups: NewPermissionGroupRepository(db),
		UserPermissions:  NewUserPermissionRepository(db),
		Users:            NewUserRepository(db),
		Wallets:          NewWalletRepository(db),
		Banks:            banks.New(db),
		CurrencyStates:   currency_states.New(db),
	}
}
