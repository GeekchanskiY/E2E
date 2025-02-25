package frontend

import (
	"fmt"
	"html/template"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("frontend.login")
	user, ok := r.Context().Value("user").(string)
	if ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println(user)
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

func (c *Controller) LoginForm(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("frontend.login.form")
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
