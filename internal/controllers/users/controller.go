package users

import (
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/operationGroups"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"

	"go.uber.org/zap"
)

type Controller struct {
	logger              *zap.Logger
	userRepo            *users.Repository
	permissionGroupRepo permission_groups.Repository
	userPermissionRepo  *user_permissions.Repository
	walletRepo          *wallets.Repository
	bankRepo            *banks.Repository
	operationGroupRepo  *operationGroups.Repository
	operationsRepo      *operations.Repository
}

func New(
	logger *zap.Logger,
	userRepo *users.Repository,
	permissionGroupRepo permission_groups.Repository,
	userPermissionRepo *user_permissions.Repository,
	walletRepo *wallets.Repository,
	bankRepo *banks.Repository,
	operationGroupRepo *operationGroups.Repository,
	operationsRepo *operations.Repository,
) *Controller {
	return &Controller{
		logger:              logger,
		userRepo:            userRepo,
		permissionGroupRepo: permissionGroupRepo,
		userPermissionRepo:  userPermissionRepo,
		walletRepo:          walletRepo,
		bankRepo:            bankRepo,
		operationsRepo:      operationsRepo,
		operationGroupRepo:  operationGroupRepo,
	}
}
