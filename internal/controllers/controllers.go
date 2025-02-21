package controllers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/users"
	"finworker/internal/repositories"
)

type Controllers struct {
	users *users.UserController
}

func NewControllers(logger *zap.Logger, repo *repositories.Repositories) *Controllers {
	userController := users.New(
		logger,
		repo.GetUsers(),
		repo.GetPermissionGroups(),
		repo.GetUserPermissions(),
		repo.GetWallets(),
		repo.GetBanks(),
	)
	return &Controllers{
		users: userController,
	}
}

func (c *Controllers) GetUsers() *users.UserController {
	return c.users
}
