package handlers

import (
	"dondarrion91/unit-converter/internal/http/utils"
	"dondarrion91/unit-converter/internal/pkg"
	"net/http"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	page := pkg.Page{
		Title: "length",
		Units: []pkg.UnitDto{
			{Name: "millimeter"},
			{Name: "centimeter"},
			{Name: "meter"},
			{Name: "kilometer"},
			{Name: "inch"},
			{Name: "foot"},
			{Name: "yard"},
			{Name: "mile"},
		},
	}

	utils.RenderTemplate(w, r, &page)
}
