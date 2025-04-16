package finance

import (
	"context"
	"html/template"
	"math"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *controller) Wallet(ctx context.Context, walletID int64) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.user.controller", zap.String("event", "got request"))

	data := utils.BuildDefaultDataMapFromContext(ctx)

	walletData, err := c.walletsRepo.Get(ctx, walletID)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["wallet"] = walletData

	distributors, err := c.distributorsRepo.GetForWallet(ctx, walletData.ID)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["distributors"] = distributors

	operationGroups, err := c.operationGroupsRepo.GetByWallet(ctx, walletData.ID)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["operationGroups"] = operationGroups

	operations, err := c.operationsRepo.GetForWallet(ctx, walletData.ID)
	if err != nil {
		c.logger.Error("frontend.wallet.controller", zap.Error(err))
		return nil, nil, err
	}

	data["operations"] = operations

	var balance float64 = 0
	for _, operation := range operations {
		balance += operation.Amount
	}
	balance = math.Round(balance*100) / 100
	data["balance"] = balance

	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.WalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	return html, data, nil
}
