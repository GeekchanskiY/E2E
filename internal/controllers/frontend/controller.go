package frontend

import (
	"embed"

	"finworker/internal/controllers/frontend/base"
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currency_states"
	"finworker/internal/repositories/distributors"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/operaton_groups"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger

	base base.Controller

	userRepo             *users.Repository
	banksRepo            *banks.Repository
	distributorsRepo     *distributors.Repository
	permissionGroupsRepo *permission_groups.Repository
	currencyStatesRepo   *currency_states.Repository
	userPermissionsRepo  *user_permissions.Repository
	walletsRepo          *wallets.Repository
	operationsRepo       *operations.Repository
	operationGroupsRepo  *operaton_groups.Repository

	secret string

	fs embed.FS
}

func New(
	logger *zap.Logger,
	userRepo *users.Repository,
	banksRepo *banks.Repository,
	distributorsRepo *distributors.Repository,
	permissionGroupsRepo *permission_groups.Repository,
	currencyStatesRepo *currency_states.Repository,
	userPermissionsRepo *user_permissions.Repository,
	walletsRepo *wallets.Repository,
	operationsRepo *operations.Repository,
	operationGroupsRepo *operaton_groups.Repository,
	secret string,
) *Controller {
	baseController := base.New(logger, userRepo, banksRepo, distributorsRepo, permissionGroupsRepo, currencyStatesRepo, userPermissionsRepo, walletsRepo, operationsRepo, operationGroupsRepo, secret)
	return &Controller{
		logger: logger,

		base: baseController,

		userRepo:             userRepo,
		banksRepo:            banksRepo,
		distributorsRepo:     distributorsRepo,
		permissionGroupsRepo: permissionGroupsRepo,
		currencyStatesRepo:   currencyStatesRepo,
		userPermissionsRepo:  userPermissionsRepo,
		walletsRepo:          walletsRepo,
		operationsRepo:       operationsRepo,
		operationGroupsRepo:  operationGroupsRepo,

		secret: secret,

		fs: templates.Fs,
	}
}

func (c *Controller) Base() base.Controller {
	return c.base
}
