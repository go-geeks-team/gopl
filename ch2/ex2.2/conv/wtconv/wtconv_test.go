package wtconv

import (
	"testing"
)

func TestLbToKg(t *testing.T) {
	weight := 10.0
	checked := Kilogram(weight / 2.205)
	unchecked := LbToKg(Pound(weight))

	t.Logf("LbToKg(%v) = %v\n", unchecked, checked)
	if checked != unchecked {
		t.Errorf("%v not equal %v\n", unchecked, checked)
	}
}

func TestKgToLb(t *testing.T) {
	weight := 10.0
	checked := Pound(weight * 2.205)
	unchecked := KgToLb(Kilogram(weight))

	t.Logf("KgToLb(%v) = %v\n", unchecked, checked)
	if checked != unchecked {
		t.Errorf("%v not equal %v\n", unchecked, checked)
	}
}
