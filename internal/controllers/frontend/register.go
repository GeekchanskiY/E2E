package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
)

func (c *Controller) Register(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.register")
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.RegisterTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
