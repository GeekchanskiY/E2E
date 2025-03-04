package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
)

func (c *Controller) Finance(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.finance")
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.FinanceTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
