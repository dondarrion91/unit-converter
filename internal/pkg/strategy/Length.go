package strategy

import "fmt"

type LengthStrategy struct{}

var factorToMM = map[string]float64{
	"millimeter": 1,
	"centimeter": 10,
	"meter":      1000,
	"kilometer":  1000 * 1000,
	"inch":       25.4,
	"foot":       304.8,
	"yard":       914.4,
	"mile":       1609344,
}

func toMM(value float64, from string) (float64, error) {
	f, ok := factorToMM[from]
	if !ok {
		return 0, fmt.Errorf("unknown length unit (from): %s", from)
	}
	return value * f, nil
}

func fromMM(valueMM float64, to string) (float64, error) {
	f, ok := factorToMM[to]
	if !ok {
		return 0, fmt.Errorf("unknown length unit (to): %s", to)
	}
	return valueMM / f, nil
}

func (l *LengthStrategy) execute(ctx *Context) float64 {
	value := float64(ctx.data.Value)

	if ctx.data.From == ctx.data.To {
		return value
	}

	mm, err := toMM(value, ctx.data.From)
	if err != nil {
		return value
	}

	out, err := fromMM(mm, ctx.data.To)
	if err != nil {
		return value
	}

	return out
}
