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

func New(logger *zap.Logger, cfg Config, repo *repositories.Repositories) *Controllers {
	userController := users.New(
		logger,
		repo.GetUsers(),
		repo.GetPermissionGroups(),
		repo.GetUserPermissions(),
		repo.GetWallets(),
		repo.GetBanks(),
		repo.GetOperationGroups(),
		repo.GetOperations(),
	)

	frontendController := frontend.New(
		logger,
		repo.GetUsers(),
		repo.GetBanks(),
		repo.GetDistributors(),
		repo.GetPermissionGroups(),
		repo.GetCurrencyStates(),
		repo.GetUserPermissions(),
		repo.GetWallets(),
		repo.GetOperations(),
		repo.GetOperationGroups(),
		cfg.GetSecret(),
	)

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
