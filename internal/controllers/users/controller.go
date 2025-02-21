package users

import (
	"go.uber.org/zap"

	"finworker/internal/repository"
)

type UserController struct {
	logger              *zap.Logger
	userRepo            *repository.UserRepository
	permissionGroupRepo *repository.PermissionGroupRepository
	userPermissionRepo  *repository.UserPermissionRepository
	walletRepo          *repository.WalletRepository
}

func New(
	logger *zap.Logger,
	userRepo *repository.UserRepository,
	permissionGroupRepo *repository.PermissionGroupRepository,
	userPermissionRepo *repository.UserPermissionRepository,
	walletRepo *repository.WalletRepository,
) *UserController {
	return &UserController{
		logger:              logger,
		userRepo:            userRepo,
		permissionGroupRepo: permissionGroupRepo,
		userPermissionRepo:  userPermissionRepo,
		walletRepo:          walletRepo,
	}
}
