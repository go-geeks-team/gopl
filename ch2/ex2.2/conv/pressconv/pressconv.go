package pressconv

import "fmt"

type Bar float64
type Pascal float64
type Atmosphere float64 // Standard atmosphere

func (b Bar) String() string {
	return fmt.Sprintf("%.6f Bar", b)
}

func (p Pascal) String() string {
	return fmt.Sprintf("%.2f Pa", p)
}

func (atm Atmosphere) String() string {
	return fmt.Sprintf("%.6f atm", atm)
}

// BarToPascal convert Bar to Pascal
func BarToPascal(b Bar) Pascal {
	return Pascal(b * 100000)
}

// BarToAtm convert Bar to "Standard atmosphere"
func BarToAtm(b Bar) Atmosphere {
	return Atmosphere(b / 1.013)
}

// PascalToBar  convert Pascal to Bar
func PascalToBar(p Pascal) Bar {
	return Bar(p / 100000)
}

// PascalToAtm convert Pascal to "Standard atmosphere"
func PascalToAtm(p Pascal) Atmosphere {
	return Atmosphere(p / 101325)
}

// AtmToBar multiplication "Standard atmosphere" to "Bar"
func AtmToBar(atm Atmosphere) Bar {
	return Bar(atm * 1.013)
}

// AtmToPascal convert "Standard atmosphere" to Pascal
func AtmToPascal(atm Atmosphere) Pascal {
	return Pascal(atm * 101300)
}
