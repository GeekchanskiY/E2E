package finance

import (
	"context"
	"embed"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/models"
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
	Finance(ctx context.Context) (*template.Template, map[string]any, error)
	CreateWallet(ctx context.Context) (*template.Template, map[string]any, error)
	CreateWalletForm(ctx context.Context, walletData models.WalletExtended) (*template.Template, map[string]any, error)
	CreateOperationGroup(ctx context.Context) (*template.Template, map[string]any, error)
	CreateOperationGroupForm(ctx context.Context, operationGroup models.OperationGroup) (*template.Template, map[string]any, error)
	CreateOperation(ctx context.Context, walletId int64) (*template.Template, map[string]any, error)
	CreateOperationForm(ctx context.Context, operation models.Operation, walletId int64) (*template.Template, map[string]any, error)
	CreateDistributor(ctx context.Context) (*template.Template, map[string]any, error)
	CreateDistributorForm(ctx context.Context, distributor models.Distributor) (*template.Template, map[string]any, error)
	Wallet(ctx context.Context, walletId int64) (*template.Template, map[string]any, error)
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
