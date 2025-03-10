package frontend

import (
	"context"
	"html/template"

	templateUtils "finworker/internal/controllers/frontend/utils"
	"finworker/internal/models"
	"finworker/internal/templates"
	"go.uber.org/zap"
)

func (c *Controller) CreateWallet(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}

func (c *Controller) CreateWalletForm(ctx context.Context, walletData models.WalletExtended) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.create_wallet.controller.form", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.CreateWalletTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := templateUtils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}
