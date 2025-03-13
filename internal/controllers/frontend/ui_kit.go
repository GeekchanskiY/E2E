package frontend

import (
	"context"
	"html/template"

	"go.uber.org/zap"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
)

func (c *Controller) UIKit(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.ui_kit.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.UIKitTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}
