package docs

import (
	"html/template"
	"path"
)

var tmpl *template.Template

func GetTemplate() *template.Template {
	if tmpl == nil {
		tmpl = template.Must(template.ParseFiles(path.Join("docs", "doc.html")))
	}

	return tmpl
}
