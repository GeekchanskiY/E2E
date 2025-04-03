package templates

import (
	"embed"
)

//go:embed *
var Fs embed.FS

const (
	BaseTemplate                     = "base.gohtml"
	FinanceTemplate                  = "finance.gohtml"
	IndexTemplate                    = "index.gohtml"
	LoginTemplate                    = "login.gohtml"
	RegisterTemplate                 = "register.gohtml"
	UserTemplate                     = "user.gohtml"
	WalletTemplate                   = "wallet.gohtml"
	CreateWalletTemplate             = "create_wallet.gohtml"
	CreateDistributorTemplate        = "create_distributor.gohtml"
	CreateOperationTemplate          = "create_operation.gohtml"
	CreateOperationGroupTemplate     = "create_operation_group.gohtml"
	UIKitTemplate                    = "ui_kit.gohtml"
	PageNotFoundTemplate             = "page_not_found.gohtml"
	FaqTemplate                      = "faq.gohtml"
	CreateWorkTemplate               = "create_work.gohtml"
	WorkTimeTemplate                 = "work_time.gohtml"
	AddUserToPermissionGroupTemplate = "add_user_to_permission_group.gohtml"
	CreatePermissionGroupTemplate    = "create_permission_group.gohtml"
	PermissionGroupsTemplate         = "permission_groups.gohtml"
)
