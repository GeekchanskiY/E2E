package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
	"go.uber.org/zap"
)

func (c *Controller) User(_ context.Context, username string) (*template.Template, error) {
	c.logger.Debug("frontend.user.controller", zap.String("event", "got request"))

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.UserTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
