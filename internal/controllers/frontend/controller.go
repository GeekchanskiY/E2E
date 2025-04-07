package frontend

import (
	"embed"

	"finworker/internal/controllers/frontend/base"
	"finworker/internal/controllers/frontend/finance"
	"finworker/internal/controllers/frontend/permissions"
	"finworker/internal/controllers/frontend/work"
	"finworker/internal/repositories/banks"
	"finworker/internal/repositories/currency_states"
	"finworker/internal/repositories/distributors"
	"finworker/internal/repositories/operations"
	"finworker/internal/repositories/operaton_groups"
	"finworker/internal/repositories/permission_groups"
	"finworker/internal/repositories/user_permissions"
	"finworker/internal/repositories/users"
	"finworker/internal/repositories/wallets"
	"finworker/internal/repositories/works"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

type Controllers interface {
	Finance() finance.Controller
	Base() base.Controller
	Work() work.Controller
	Permissions() permissions.Controller
}

type controllers struct {
	logger *zap.Logger

	base        base.Controller
	finance     finance.Controller
	work        work.Controller
	permissions permissions.Controller

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
	workRepo *works.Repository,
	secret string,
) Controllers {
	baseController := base.New(logger, userRepo, banksRepo, distributorsRepo, permissionGroupsRepo, currencyStatesRepo, userPermissionsRepo, walletsRepo, operationsRepo, operationGroupsRepo, secret)
	financeController := finance.New(logger, userRepo, banksRepo, distributorsRepo, permissionGroupsRepo, currencyStatesRepo, userPermissionsRepo, walletsRepo, operationsRepo, operationGroupsRepo, secret)
	workController := work.New(logger, userRepo, workRepo, secret)
	permissionsController := permissions.New(logger, userRepo, permissionGroupsRepo, userPermissionsRepo)

	return &controllers{
		logger: logger,

		base:        baseController,
		finance:     financeController,
		work:        workController,
		permissions: permissionsController,

		secret: secret,

		fs: templates.Fs,
	}
}

func (c *controllers) Base() base.Controller {
	return c.base
}

func (c *controllers) Finance() finance.Controller {
	return c.finance
}

func (c *controllers) Work() work.Controller {
	return c.work
}

func (c *controllers) Permissions() permissions.Controller { return c.permissions }
