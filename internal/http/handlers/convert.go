package handlers

import (
	"dondarrion91/unit-converter/internal/pkg"
	"encoding/json"
	"net/http"
)

type ConvertResponse struct {
	Value          int     `json:"value"`
	From           string  `json:"from"`
	To             string  `json:"to"`
	UnitType       string  `json:"unitType"`
	ConvertedValue float64 `json:"convertedValue"`
}

func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload ConvertResponse

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	unit, err := pkg.GetUnit(payload.UnitType, payload.Value, payload.From, payload.To)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ConvertResponse{
		Value:          unit.GetValue(),
		From:           unit.GetFrom(),
		To:             unit.GetTo(),
		UnitType:       unit.GetName(),
		ConvertedValue: unit.GetConvertedValue(),
	})
}
