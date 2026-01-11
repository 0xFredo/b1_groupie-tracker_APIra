package utils

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// PageData holds common data for all pages
type PageData struct {
	Title           string
	ActiveTab       string
	SearchQuery     string
	SearchExpanded  bool
	ContentTemplate string
	Data            interface{}
}

// InitTemplates loads all HTML templates
func InitTemplates() error {
	var err error
	templates, err = template.New("").Funcs(template.FuncMap{
		"json": func(data interface{}) (string, error) {
			bytes, err := json.Marshal(data)
			if err != nil {
				return "", err
			}
			return string(bytes), nil
		},
	}).ParseGlob(filepath.Join("web", "templates", "*.html"))
	return err
}

// RenderTemplate renders a template with data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return templates.ExecuteTemplate(w, tmpl, data)
}

// Chargement/exécution templates
