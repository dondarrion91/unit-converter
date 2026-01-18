package strategy

import (
	"fmt"
	"strings"
)

type TemperatureStrategy struct{}

func toKelvin(value float64, from string) (float64, error) {
	switch strings.ToLower(from) {
	case "celsius":
		return value + 273.15, nil
	case "kelvin":
		return value, nil
	case "fahrenheit":
		return (value-32)*5/9 + 273.15, nil
	default:
		return 0, fmt.Errorf("unknown temperature unit (from): %s", from)
	}
}

func fromKelvin(valueK float64, to string) (float64, error) {
	switch strings.ToLower(to) {
	case "celsius":
		return valueK - 273.15, nil
	case "kelvin":
		return valueK, nil
	case "fahrenheit":
		return (valueK-273.15)*9/5 + 32, nil
	default:
		return 0, fmt.Errorf("unknown temperature unit (to): %s", to)
	}
}

func (t *TemperatureStrategy) execute(ctx *Context) float64 {
	value := float64(ctx.data.Value)

	if ctx.data.From == ctx.data.To {
		return value
	}

	k, err := toKelvin(value, ctx.data.From)
	if err != nil {
		return value
	}

	out, err := fromKelvin(k, ctx.data.To)
	if err != nil {
		return value
	}

	return out
}
