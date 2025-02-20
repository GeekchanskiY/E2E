package users

import (
	"go.uber.org/zap"

	"finworker/internal/repository"
)

type UserController struct {
	logger              *zap.Logger
	userRepo            *repository.UserRepository
	permissionGroupRepo *repository.PermissionGroupRepository
}

func New(
	logger *zap.Logger,
	userRepo *repository.UserRepository,
	permissionGroupRepo *repository.PermissionGroupRepository,
) *UserController {
	return &UserController{
		logger:              logger,
		userRepo:            userRepo,
		permissionGroupRepo: permissionGroupRepo,
	}
}
