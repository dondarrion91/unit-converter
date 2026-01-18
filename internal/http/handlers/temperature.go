package handlers

import (
	"dondarrion91/unit-converter/internal/http/utils"
	"dondarrion91/unit-converter/internal/pkg"
	"net/http"
)

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	page := pkg.Page{
		Title: "temperature",
		Units: []pkg.UnitDto{
			{Name: "Celsius"},
			{Name: "Fahrenheit"},
			{Name: "Kelvin"},
		},
	}

	utils.RenderTemplate(w, r, &page)
}
