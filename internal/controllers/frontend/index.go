package frontend

import (
	"context"
	"html/template"

	"finworker/internal/templates"
)

func (c *Controller) Index(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.index")
	html, err := template.ParseFS(c.fs, templates.BaseTemplate, templates.IndexTemplate)
	if err != nil {
		return nil, err
	}

	return html, nil
}
