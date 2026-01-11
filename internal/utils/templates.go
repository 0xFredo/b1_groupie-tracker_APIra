package utils

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var templates *template.Template

type PageData struct {
	Title           string
	ActiveTab       string
	SearchQuery     string
	SearchExpanded  bool
	ContentTemplate string
	Data            interface{}
}

func formatLocation(loc string) string {
	loc = strings.ReplaceAll(loc, "_", " ")
	parts := strings.Split(loc, "-")
	for i, part := range parts {
		parts[i] = strings.Title(strings.TrimSpace(part))
	}
	return strings.Join(parts, ", ")
}

func formatDate(dateStr string) string {
	parts := strings.Split(dateStr, "-")
	if len(parts) != 3 {
		return dateStr
	}
	day, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	year, _ := strconv.Atoi(parts[2])

	if month < 1 || month > 12 {
		return dateStr
	}

	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Format("2 January 2006")
}

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
		"formatLocation": formatLocation,
		"formatDate":     formatDate,
	}).ParseGlob(filepath.Join("web", "templates", "*.html"))
	return err
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	return templates.ExecuteTemplate(w, tmpl, data)
}
