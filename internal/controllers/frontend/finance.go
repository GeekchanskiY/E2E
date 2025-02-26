package frontend

import (
	"context"
	"html/template"
)

func (c *Controller) Finance(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.finance")
	html, err := template.ParseFS(c.fs, "base.gohtml", "finance.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}
