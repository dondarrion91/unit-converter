package router

import (
	"dondarrion91/unit-converter/internal/http/handlers"
	"net/http"
)

func Router() {
	http.HandleFunc("/length/", handlers.LengthHandler)
	http.HandleFunc("/weight/", handlers.WeightHandler)
	http.HandleFunc("/temperature/", handlers.TemperatureHandler)
	http.HandleFunc("/convert", handlers.ConvertHandler)
}
