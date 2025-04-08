package htmlPages

import (
	"html/template"
	"net/http"

	"cyoa/entity"
)

func buildPage(arc *entity.Arc) *template.Template {
	tmpl := template.Must(template.ParseFiles("res/template.html"))
	return tmpl
}

func NewHandler(arc *entity.Arc) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		err := buildPage(arc).Execute(w, arc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
