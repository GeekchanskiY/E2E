package frontend

import (
	"html/template"
	"net/http"

	"finworker/internal/models/templates"
)

func (c *Controller) Index(w http.ResponseWriter, _ *http.Request) {
	c.logger.Info("frontend.index")
	html, err := template.ParseFS(c.fs, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, templates.Index{Text: "Hello World!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
