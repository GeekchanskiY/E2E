package templates

import (
	"embed"
)

//go:embed *
var Fs embed.FS

const (

	// base templates

	BaseTemplate         = "base/base.gohtml"
	IndexTemplate        = "base/index.gohtml"
	FaqTemplate          = "base/faq.gohtml"
	UIKitTemplate        = "base/ui_kit.gohtml"
	LoginTemplate        = "base/login.gohtml"
	RegisterTemplate     = "base/register.gohtml"
	UserTemplate         = "base/user.gohtml"
	PageNotFoundTemplate = "base/page_not_found.gohtml"

	// finance templates

	FinanceTemplate              = "finance/finance.gohtml"
	WalletTemplate               = "finance/wallet.gohtml"
	CreateWalletTemplate         = "finance/create_wallet.gohtml"
	CreateDistributorTemplate    = "finance/create_distributor.gohtml"
	CreateOperationTemplate      = "finance/create_operation.gohtml"
	CreateOperationGroupTemplate = "finance/create_operation_group.gohtml"

	// work templates

	CreateWorkTemplate = "work/create_work.gohtml"
	WorkTimeTemplate   = "work/work_time.gohtml"

	// permission groups templates

	AddUserToPermissionGroupTemplate = "permissions/add_user_to_permission_group.gohtml"
	CreatePermissionGroupTemplate    = "permissions/create_permission_group.gohtml"
	PermissionGroupsTemplate         = "permissions/permission_groups.gohtml"
	PermissionGroupTemplate          = "permissions/permission_group.gohtml"
)
