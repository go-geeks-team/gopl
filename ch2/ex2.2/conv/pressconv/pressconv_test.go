package pressconv

import (
	"testing"
)

// Bar to Pascal
func TestBarToPascal(t *testing.T) {
	bar := Bar(10.0)
	checked := Pascal(bar * 100000)
	unchecked := BarToPascal(bar)

	if unchecked != checked {
		t.Errorf("BarToPascal(%v) not equal %v", unchecked, checked)
	}
}

// Bar to Atmosphere
func TestBarToAtm(t *testing.T) {
	bar := Bar(10.0)
	checked := Atmosphere(bar / 1.013)
	unchecked := BarToAtm(bar)

	if unchecked != checked {
		t.Errorf("BarToAtm(%v) not equal %v", unchecked, checked)
	}
}

// PascalToBar
func TestPascalToBar(t *testing.T) {
	pas := Pascal(100000)
	checked := Bar(pas / 100000)
	unchecked := PascalToBar(pas)

	if unchecked != checked {
		t.Errorf("PascalToBar(%v) not equal %v", unchecked, checked)
	}
}

func TestPascalToAtm(t *testing.T) {
	pas := Pascal(1)
	checked := Atmosphere(pas / 101325)
	unchecked := PascalToAtm(pas)

	if unchecked != checked {
		t.Errorf("PascalToAtm(%v) not equal %v", unchecked, checked)
	}
}
