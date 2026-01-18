package strategy

import "fmt"

type WeightStrategy struct{}

var factorTomg = map[string]float64{
	"milligram": 1,            // base
	"gram":      1000,         // 1 g = 1000 mg
	"kilogram":  1000000,      // 1 kg = 1e6 mg
	"ounce":     28349.523125, // 1 oz = 28.349523125 g
	"pound":     453592.37,    // 1 lb = 453.59237 g
}

func tomg(value float64, from string) (float64, error) {
	f, ok := factorTomg[from]
	if !ok {
		return 0, fmt.Errorf("unknown weight unit (from): %s", from)
	}
	return value * f, nil
}

func fromg(valuemg float64, to string) (float64, error) {
	f, ok := factorTomg[to]
	if !ok {
		return 0, fmt.Errorf("unknown weight unit (to): %s", to)
	}
	return valuemg / f, nil
}

func (l *WeightStrategy) execute(ctx *Context) float64 {
	value := float64(ctx.data.Value)

	if ctx.data.From == ctx.data.To {
		return value
	}

	mm, err := tomg(value, ctx.data.From)
	if err != nil {
		return value
	}

	out, err := fromg(mm, ctx.data.To)
	if err != nil {
		return value
	}

	return out
}
