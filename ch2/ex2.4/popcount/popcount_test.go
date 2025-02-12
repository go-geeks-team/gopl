package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	x := byte(255) // uint8(...) equal byte(...) (is the set of all unsigned 8-bit integers. Range: 0 through 255.)
	u := uint64(x)

	t.Logf("   PopCount(%b) Количество установленных битов - %d\n", u, PopCount(u))
	t.Logf("PopCountNew(%b) Количество установленных битов - %d\n", u, NewPopCount(u))

	if PopCount(u) != NewPopCount(u) {
		t.Errorf("PopCount(%[1]d) not equal NewPopCount(%[1]d\n", u)
	}
}
