package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
)

func (c *Controller) User(_ context.Context, username string) (*template.Template, error) {
	c.logger.Info("frontend.user")

	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.UserTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
