package pkg

import (
	"fmt"
)

func GetUnit(unitType string, value int, from string, to string) (IUnit, error) {
	if unitType == "length" {
		return NewLength(value, from, to), nil
	}

	if unitType == "weight" {
		return NewWeight(value, from, to), nil
	}

	if unitType == "temperature" {
		return NewTemperature(value, from, to), nil
	}

	return nil, fmt.Errorf("Wrong unit type passed")
}
