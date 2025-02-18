package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Distributors     *DistributorsRepository
	OperationGroups  *OperationGroupRepository
	Operations       *OperationRepository
	PermissionGroups *PermissionGroupRepository
	UserPermissions  *UserPermissionRepository
	Users            *UserRepository
	Wallets          *WalletRepository
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
	}
}
