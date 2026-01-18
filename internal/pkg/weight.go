package pkg

import "dondarrion91/unit-converter/internal/pkg/strategy"

type Weight struct {
	Unit
}

func getShortWeightUnit(unit string) string {
	switch unit {
	case "milligram":
		return "mg"
	case "gram":
		return "g"
	case "kilogram":
		return "kg"
	case "ounce":
		return "oz"
	case "pound":
		return "lb"
	default:
		return unit
	}
}

func NewWeight(value int, from string, to string) IUnit {
	weight := &Weight{
		Unit: Unit{
			Name:  "weight",
			Value: value,
			From:  getShortWeightUnit(from),
			To:    getShortWeightUnit(to),
		},
	}

	str := &strategy.WeightStrategy{}
	convertedValue := weight.Convert(str, value, from, to)
	weight.SetConvertedValue(convertedValue)

	return weight
}
