package tempconv

import (
	"testing"
)

// Celsius to Fahrenheit
// Formula: (0°C × 9/5) + 32
func TestCToF(t *testing.T) {
	temp := 0
	checked := Fahrenheit((temp * 9 / 5) + 32) // correct formula
	unchecked := CToF(Celsius(temp))

	t.Logf("CToF(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("CToF(%v) not equal %v\n", unchecked, checked)
	}
}

// Fahrenheit to Celsius
// Formula: (0°F − 32) × 5/9
func TestFToC(t *testing.T) {
	temp := 212.0
	checked := Celsius((temp - 32) * 5 / 9) // correct formula
	unchecked := FToC(Fahrenheit(temp))

	t.Logf("FToC(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("FToC(%v) not equal %v\n", unchecked, checked)
	}
}

// Kelvin to Fahrenheit
// Formula: 0K − 273.15) × 9/5 + 32
func TestKelToFah(t *testing.T) {
	temp := 0.0
	checked := Fahrenheit((temp-273.15)*9/5 + 32) // correct formula
	unchecked := KToF(Kelvin(temp))

	t.Logf("KToF(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("KToF(%v) not equal %v\n", unchecked, checked)
	}

}

// Kelvin to Celsius
// Formula: 0K − 273.15 = -273.1°C
func TestKelToCel(t *testing.T) {
	temp := 373.15
	checked := Celsius(temp - 273.15) // correct formula
	unchecked := KToC(Kelvin(temp))

	t.Logf("KToC(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("KToC(%v) not equal %v", unchecked, checked)
	}
}
