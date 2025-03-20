package frontend

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *Controller) Wallet(ctx context.Context, walletId int) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.user.controller", zap.String("event", "got request"))

	data := utils.BuildDefaultDataMapFromContext(ctx)

	walletData, err := c.walletsRepo.Get(ctx, walletId)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["wallet"] = walletData

	distributors, err := c.distributorsRepo.GetForWallet(ctx, walletData.Id)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["distributors"] = distributors

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.WalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	return html, data, nil
}
