package web

import "net/http"

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type indexTemplate struct {
			Page        string
			Title       string
			Description string
			Version     string
		}

		renderTemplate(w, "index.html", indexTemplate{
			Page:        "Home",
			Title:       "Gophers Management System",
			Description: "Gophers is a simple CRUD application to manage gophers.",
			Version:     "v1.0.0",
		})
	}
}
