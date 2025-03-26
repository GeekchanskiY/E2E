package utils

import (
	"embed"
	"fmt"
	"html/template"
	"strings"
	"time"
)

// GenerateTemplate generates template, and fills it with all template functions
func GenerateTemplate(fs embed.FS, baseTemplate, targetTemplate string) (*template.Template, error) {
	html, err := template.New("test").Funcs(template.FuncMap{
		"formatFloat": formatFloat,
		"formatTime":  formatTime,
	}).ParseFS(fs, baseTemplate, targetTemplate)

	return html, err
}

// formatFloat is a template func, which converts floats to readable state
func formatFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// formatTime is a template func, which converts time to readable state
func formatTime(t time.Time) string {
	return strings.Replace(t.Format("02 mnt 2006 15:04:05"), "mnt", getMonthRuAbbr(t.Month()), 1)
}

func getMonthRuAbbr(m time.Month) string {
	switch m {
	case time.January:
		return "янв"
	case time.February:
		return "фев"
	case time.March:
		return "мар"
	case time.April:
		return "апр"
	case time.May:
		return "май"
	case time.June:
		return "июн"
	case time.July:
		return "июл"
	case time.August:
		return "авг"
	case time.September:
		return "сен"
	case time.October:
		return "окт"
	case time.November:
		return "ноя"
	case time.December:
		return "дек"
	}

	return ""
}
