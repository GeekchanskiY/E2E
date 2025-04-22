package finance

import (
	"context"
	"html/template"

	"finworker/internal/config"
	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"

	"go.uber.org/zap"
)

func (c *controller) Finance(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.finance.controller", zap.String("event", "got request"))

	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.FinanceTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	wallets, err := c.walletsRepo.GetByUsername(ctx, ctx.Value(config.UsernameContextKey).(string))
	if err != nil {
		return nil, nil, err
	}

	data["wallets"] = wallets

	return html, data, nil
}
