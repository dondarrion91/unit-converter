package pkg

import "dondarrion91/unit-converter/internal/pkg/strategy"

type Length struct {
	Unit
}

func getShortUnit(unit string) string {
	switch unit {
	case "millimeter":
		return "mm"
	case "centimeter":
		return "cm"
	case "meter":
		return "m"
	case "kilometer":
		return "km"
	case "inch":
		return "in"
	case "foot":
		return "ft"
	case "yard":
		return "yd"
	case "mile":
		return "mi"
	default:
		return unit
	}
}

func NewLength(value int, from string, to string) IUnit {
	length := &Length{
		Unit: Unit{
			Name:  "length",
			Value: value,
			From:  getShortUnit(from),
			To:    getShortUnit(to),
		},
	}

	str := &strategy.LengthStrategy{}
	convertedValue := length.Convert(str, value, from, to)
	length.SetConvertedValue(convertedValue)

	return length
}
