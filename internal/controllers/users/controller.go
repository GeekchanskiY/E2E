package users

import (
	"finworker/internal/repositories/banks"
	"go.uber.org/zap"

	"finworker/internal/repositories"
)

type UserController struct {
	logger              *zap.Logger
	userRepo            *repositories.UserRepository
	permissionGroupRepo *repositories.PermissionGroupRepository
	userPermissionRepo  *repositories.UserPermissionRepository
	walletRepo          *repositories.WalletRepository
	bankRepo            *banks.Repository
}

func New(
	logger *zap.Logger,
	userRepo *repositories.UserRepository,
	permissionGroupRepo *repositories.PermissionGroupRepository,
	userPermissionRepo *repositories.UserPermissionRepository,
	walletRepo *repositories.WalletRepository,
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
