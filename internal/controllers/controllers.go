package controllers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/users"
	"finworker/internal/repository"
)

type Controllers struct {
	users *users.UserController
}

func NewControllers(logger *zap.Logger, repo *repository.Repositories) *Controllers {
	userController := users.New(logger, repo.Users, repo.PermissionGroups, repo.UserPermissions)
	return &Controllers{
		users: userController,
	}
}

func (c *Controllers) GetUsers() *users.UserController {
	return c.users
}
