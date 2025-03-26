package utils

import (
	"embed"
	"fmt"
	"html/template"
)

// GenerateTemplate generates template, and fills it with all template functions
func GenerateTemplate(fs embed.FS, baseTemplate, targetTemplate string) (*template.Template, error) {
	html, err := template.New("test").Funcs(template.FuncMap{
		"formatFloat": formatFloat,
	}).ParseFS(fs, baseTemplate, targetTemplate)

	return html, err
}

// formatFloat is a template func, which converts floats to readable state
func formatFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
