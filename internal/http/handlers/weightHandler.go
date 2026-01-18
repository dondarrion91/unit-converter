package handlers

import (
	"dondarrion91/unit-converter/internal/http/utils"
	"dondarrion91/unit-converter/internal/pkg"
	"net/http"
)

func WeightHandler(w http.ResponseWriter, r *http.Request) {
	weight := pkg.Page{
		Title: "weight",
		Units: []pkg.UnitDto{
			{Name: "milligram"},
			{Name: "gram"},
			{Name: "kilogram"},
			{Name: "ounce"},
			{Name: "poun"},
		},
	}

	utils.RenderTemplate(w, r, &weight)
}
