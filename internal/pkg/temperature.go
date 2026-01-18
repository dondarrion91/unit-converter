package pkg

import (
	"dondarrion91/unit-converter/internal/pkg/strategy"
	"strings"
)

type Temperature struct {
	Unit
}

func getShortTemperatureUnit(unit string) string {
	switch strings.ToLower(unit) {
	case "celsius":
		return "°C"
	case "kelvin":
		return "K"
	case "fahrenheit":
		return "°F"
	default:
		return unit
	}
}

func NewTemperature(value int, from string, to string) IUnit {
	temperature := &Temperature{
		Unit: Unit{
			Name:  "temperature",
			Value: value,
			From:  getShortTemperatureUnit(from),
			To:    getShortTemperatureUnit(to),
		},
	}

	str := &strategy.TemperatureStrategy{}
	convertedValue := temperature.Convert(str, value, from, to)
	temperature.SetConvertedValue(convertedValue)

	return temperature
}
