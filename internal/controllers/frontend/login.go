package frontend

import (
	"context"
	"html/template"
)

func (c *Controller) Login(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.login")

	html, err := template.ParseFS(c.fs, "base.gohtml", "login.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}

func (c *Controller) LoginForm(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.login.form")
	html, err := template.ParseFS(c.fs, "base.gohtml", "login.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}
