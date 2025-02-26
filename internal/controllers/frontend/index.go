package frontend

import (
	"context"
	"html/template"
)

func (c *Controller) Index(_ context.Context) (*template.Template, error) {
	c.logger.Info("frontend.index")
	html, err := template.ParseFS(c.fs, "base.gohtml", "index.gohtml")
	if err != nil {
		return nil, err
	}

	return html, nil
}
