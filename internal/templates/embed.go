package templates

import (
	"embed"
)

//go:embed *
var Fs embed.FS

const (
	BaseTemplate     = "base.gohtml"
	FinanceTemplate  = "finance.gohtml"
	IndexTemplate    = "index.gohtml"
	LoginTemplate    = "login.gohtml"
	RegisterTemplate = "register.gohtml"
	UserTemplate     = "user.gohtml"
	WalletTemplate   = "wallet.gohtml"
)
