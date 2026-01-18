package pkg

import "dondarrion91/unit-converter/internal/pkg/strategy"

type IUnit interface {
	SetName(name string)
	GetName() string
	SetValue(value int)
	GetValue() int
	GetFrom() string
	SetFrom(from string)
	GetTo() string
	SetTo(to string)
	SetConvertedValue(value float64)
	GetConvertedValue() float64
	Convert(str strategy.Strategy, value int, from string, to string) float64
}

type Unit struct {
	Name           string
	Value          int
	From           string
	To             string
	ConvertedValue float64
}

func (u *Unit) SetName(name string) {
	u.Name = name
}

func (u *Unit) GetName() string {
	return u.Name
}

func (u *Unit) GetValue() int {
	return u.Value
}

func (u *Unit) SetValue(value int) {
	u.Value = value
}

func (u *Unit) GetFrom() string {
	return u.From
}

func (u *Unit) SetFrom(from string) {
	u.From = from
}

func (u *Unit) GetTo() string {
	return u.To
}

func (u *Unit) SetTo(to string) {
	u.To = to
}

func (u *Unit) GetConvertedValue() float64 {
	return u.ConvertedValue
}

func (u *Unit) SetConvertedValue(value float64) {
	u.ConvertedValue = value
}

func (u *Unit) Convert(str strategy.Strategy, value int, from string, to string) float64 {
	if from == to {
		return float64(value)
	}

	data := strategy.ContextData{
		Value: value,
		From:  from,
		To:    to,
	}

	ctx := strategy.InitContext(str, data)

	return float64(ctx.Execute())
}
