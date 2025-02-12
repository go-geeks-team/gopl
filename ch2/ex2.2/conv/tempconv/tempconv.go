// Пакет tempconv выполняет преобразование температур.

package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

// CToF преобразует температуру по Цельсию в температуру по Фаренгейту.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC преобразует температуру по Фаренгейту в температуру по Цельсию
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
