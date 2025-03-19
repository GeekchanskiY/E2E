package templates

import (
	"embed"
)

//go:embed *
var Fs embed.FS

const (
	BaseTemplate         = "base.gohtml"
	FinanceTemplate      = "finance.gohtml"
	IndexTemplate        = "index.gohtml"
	LoginTemplate        = "login.gohtml"
	RegisterTemplate     = "register.gohtml"
	UserTemplate         = "user.gohtml"
	WalletTemplate       = "wallet.gohtml"
	CreateWalletTemplate = "create_wallet.gohtml"
	UIKitTemplate        = "ui_kit.gohtml"
	PageNotFoundTemplate = "page_not_found.gohtml"
	FaqTemplate          = "faq.gohtml"
)
