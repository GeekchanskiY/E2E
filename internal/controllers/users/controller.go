package users

import (
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"

	"go.uber.org/zap"
)

type UserController struct {
	logger              *zap.Logger
	userRepo            *users.Repository
	permissionGroupRepo *permission_groups.Repository
	userPermissionRepo  *user_permissions.Repository
	walletRepo          *wallets.Repository
	bankRepo            *banks.Repository
}

func New(
	logger *zap.Logger,
	userRepo *users.Repository,
	permissionGroupRepo *permission_groups.Repository,
	userPermissionRepo *user_permissions.Repository,
	walletRepo *wallets.Repository,
	bankRepo *banks.Repository,
) *UserController {
	return &UserController{
		logger:              logger,
		userRepo:            userRepo,
		permissionGroupRepo: permissionGroupRepo,
		userPermissionRepo:  userPermissionRepo,
		walletRepo:          walletRepo,
		bankRepo:            bankRepo,
	}
}
