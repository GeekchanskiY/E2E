package frontend

import (
	"context"
	"html/template"

	"finworker/internal/controllers/frontend/utils"
	"finworker/internal/templates"
	"go.uber.org/zap"
)

func (c *Controller) Register(ctx context.Context) (*template.Template, map[string]any, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, nil, err
	}

	data := utils.BuildDefaultDataMapFromContext(ctx)

	return html, data, nil
}
