package utils

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var templates *template.Template

// InitTemplates loads all HTML templates
func InitTemplates() error {
	var err error
	templates, err = template.ParseGlob(filepath.Join("web", "templates", "*.html"))
	return err
}

// RenderTemplate renders a template with data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return templates.ExecuteTemplate(w, tmpl, data)
}

// Chargement/exécution templates
