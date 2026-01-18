package utils

import (
	"dondarrion91/unit-converter/internal/pkg"
	"html/template"
	"net/http"
)

var Layouts = template.Must(template.ParseGlob("web/templates/layout/*.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request, p *pkg.Page) {
	length := template.Must(template.Must(Layouts.Clone()).ParseFiles("web/templates/index.html"))

	err := length.ExecuteTemplate(w, "base.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
