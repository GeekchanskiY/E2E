package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
	"go.uber.org/zap"
)

func (c *Controller) Register(_ context.Context) (*template.Template, error) {
	c.logger.Debug("frontend.register.controller", zap.String("event", "got request"))
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
