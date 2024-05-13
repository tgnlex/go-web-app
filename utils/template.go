package utils

import (
	"html/template"
)

func CreateTemplate(html string) *template.Template {
	tpl := template.Must(template.ParseFiles(html))
	return tpl
}
