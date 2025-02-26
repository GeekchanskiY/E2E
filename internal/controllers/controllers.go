package controllers

import (
	"go.uber.org/zap"

	"finworker/internal/controllers/frontend"

	"finworker/internal/controllers/users"
	"finworker/internal/repositories"
)

type Controllers struct {
	users    *users.Controller
	frontend *frontend.Controller
}

func New(logger *zap.Logger, repo *repositories.Repositories) *Controllers {
	userController := users.New(
		logger,
		repo.GetUsers(),
		repo.GetPermissionGroups(),
		repo.GetUserPermissions(),
		repo.GetWallets(),
		repo.GetBanks(),
	)

	frontendController := frontend.New(logger)

	return &Controllers{
		users:    userController,
		frontend: frontendController,
	}
}

func (c *Controllers) GetUsers() *users.Controller {
	return c.users
}

func (c *Controllers) GetFrontend() *frontend.Controller {
	return c.frontend
}
