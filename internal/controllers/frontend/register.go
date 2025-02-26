package frontend

import (
	"context"
	"html/template"
)

func (c *Controller) Register(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.register")
	html, err := template.ParseFS(c.fs, "base.gohtml", "register.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}
