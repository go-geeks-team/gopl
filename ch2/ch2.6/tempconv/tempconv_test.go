package tempconv

import (
	"testing"
)

// Celsius to Fahrenheit
// Formula: (0°C × 9/5) + 32 = 33.8°F
func TestCToF(t *testing.T) {
	temp := 0.0
	checked := Fahrenheit((temp * 9 / 5) + 32) // correct formula
	unchecked := CToF(Celsius(temp))

	t.Logf("CToF(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("CToF(%v) not equal %v ", unchecked, checked)
	}
}

// Fahrenheit to Celsius
// Formula: (0°F − 32) × 5/9 = -17.22°C
func TestFToC(t *testing.T) {
	temp := 212.0
	checked := Celsius((temp - 32) * 5 / 9) // correct formula
	unchecked := FToC(Fahrenheit(temp))

	t.Logf("FToC(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("FToC(%v) not equal %v", unchecked, checked)
	}
}
