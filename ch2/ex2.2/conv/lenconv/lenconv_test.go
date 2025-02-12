package lenconv

import (
	"testing"
)

// Foot to Meter
func TestFtToMet(t *testing.T) {
	length := Foot(10.0)
	checked := Meter(length / 3.281)
	unchecked := FtToMet(length)

	t.Logf("FtToMet(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("%v not equal %v\n", unchecked, checked)
	}
}

// Meter to Foot
func TestMetToFt(t *testing.T) {
	length := 10.0
	checked := Foot(length * 3.281)
	unchecked := MetToFt(Meter(length))

	t.Logf("MetToFt(%v) = %v\n", unchecked, checked)

	if checked != unchecked {
		t.Errorf("%v not equal %v\n", unchecked, checked)
	}
}
