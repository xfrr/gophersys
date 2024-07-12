package web

import (
	"embed"
	"html/template"
)

//go:embed templates/* templates/layout/*
var templateFS embed.FS

var templates *template.Template

func (a App) loadTemplates(fm template.FuncMap) error {
	var err error
	templates, err = template.New("").
		Funcs(fm).
		ParseFS(
			templateFS,
			"templates/*.html",
			"templates/layout/*.html",
		)
	if err != nil {
		return err
	}

	return nil
}
