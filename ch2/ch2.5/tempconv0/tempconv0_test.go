package tempconv

import (
	"testing"
)

// Celsius to Fahrenheit
// Formula: (0°C × 9/5) + 32 = 33.8°F
func TestCToF(t *testing.T) {
	temp := 0.0
	checked := Fahrenheit((temp * 9 / 5) + 32)
	needCheck := CToF(Celsius(temp))

	t.Logf("CToF(%v) = %v\n", needCheck, checked)

	if checked != needCheck {
		t.Errorf("CToF(%v) not equal %v", needCheck, checked)
	}
}

// Fahrenheit to Celsius
// Formula: (0°F − 32) × 5/9 = -17.22°C
func TestFToC(t *testing.T) {
	temp := 212.0
	checked := Celsius((temp - 32) * 5 / 9)
	needCheck := FToC(Fahrenheit(temp))

	t.Logf("FToC(%v) = %v\n", needCheck, checked)

	if checked != needCheck {
		t.Errorf("FToC(%v) not equal %v", needCheck, checked)
	}
}
