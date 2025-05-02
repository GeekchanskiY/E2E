package base

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/templates"
)

func (c *controller) FAQ(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.FAQ.controller", zap.String("event", "got request"))

	html, err := utils.GenerateTemplate(c.fs, templates.BaseTemplate, templates.FaqTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}
