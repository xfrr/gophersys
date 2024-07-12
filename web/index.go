package web

import "net/http"

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type IndexTemplate struct {
			Page        string
			Title       string
			Description string
			Version     string
		}

		render(w, "index", IndexTemplate{
			Page:        "Home",
			Title:       "Gophers Management System",
			Description: "Gophers is a simple CRUD application to manage gophers.",
			Version:     "v1.0.0",
		})
	}
}

func render(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		panic(err)
	}
}
