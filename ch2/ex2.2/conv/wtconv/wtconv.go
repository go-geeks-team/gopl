package wtconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64

func (p Pound) String() string {
	return fmt.Sprintf("%.2f lb", p)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%.2f kg", kg)
}

// LbToKg Pound to Kilogram
func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb / 2.205)
}

// KgToLb Kilogram to Pound
func KgToLb(kg Kilogram) Pound {
	return Pound(kg * 2.205)
}
