package base

import (
	"context"
	"embed"
	"html/template"

	"go.uber.org/zap"

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
)

type Controller interface {
	FAQ(ctx context.Context) (*template.Template, map[string]any, error)
	Index(ctx context.Context) (*template.Template, map[string]any, error)
	Login(ctx context.Context) (*template.Template, map[string]any, error)
	LoginForm(ctx context.Context, username, password string) (*template.Template, map[string]any, string, string, error)
	PageNotFound(ctx context.Context) (*template.Template, map[string]any, error)
	Register(ctx context.Context) (*template.Template, map[string]any, error)
	RegisterForm(ctx context.Context, username, name, password, repeatPassword, gender, birthday, bank, salary, currency, payday string) (*template.Template, map[string]any, string, string, error)
	UIKit(ctx context.Context) (*template.Template, map[string]any, error)
	User(ctx context.Context, username string) (*template.Template, map[string]any, error)
}

type controller struct {
	logger *zap.Logger

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
) Controller {
	return &controller{
		logger: logger,

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
