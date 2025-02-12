package rev

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	a := [8]int{0, 1, 2, 3, 4, 5}

	reverse(&a)

	t.Logf("%v", a)

	if fmt.Sprint(a[:]) != "[0 0 5 4 3 2 1 0]" {
		t.Errorf("%v", a[:])
	}
}
