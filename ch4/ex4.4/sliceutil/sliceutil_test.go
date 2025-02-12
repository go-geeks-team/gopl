package sliceutil

import (
	"slices"
	"testing"
)

func TestRotateEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	rotateOld(s1)
	rotateNew(s2)

	if !slices.Equal(s1, s2) {
		t.Logf("rotateOld(%v) is not equal rotateNew(%v)", s1, s2)
	}
}

func TestRotateNotEqual(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3, 2, 1}

	rotateOld(s1)
	rotateNew(s2)

	if slices.Equal(s1, s2) {
		t.Errorf("rotateOld(%v) is not equal rotateNew(%v)", s1, s2)
	}
}
