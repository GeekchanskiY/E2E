package frontend

import (
	"context"
	"html/template"
)

func (c *Controller) User(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.user")

	html, err := template.ParseFS(c.fs, "base.gohtml", "user.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}
