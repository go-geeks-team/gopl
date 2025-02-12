package rev

import (
	"fmt"
	"testing"
)

func TestReverseOld(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverseOld(a[:])

	t.Logf("%v", a)
	if fmt.Sprint(a[:]) != "[5 4 3 2 1 0]" {
		t.Errorf("%v", a[:])
	}
}

func TestReverseNew(t *testing.T) {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverseNew(a[:])

	t.Logf("%v", a)
	if fmt.Sprint(a[:]) != "[5 4 3 2 1 0]" {
		t.Errorf("%v", a[:])
	}
}
