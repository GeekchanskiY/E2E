package frontend

import (
	"html/template"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, _ *http.Request) {
	c.logger.Info("frontend.login")
	html, err := template.ParseFS(c.fs, "base.gohtml", "login.gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
