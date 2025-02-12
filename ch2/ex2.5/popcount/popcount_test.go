package popcount

import (
	"math/rand"
	"testing"
)

func TestNewPopCount(t *testing.T) {
	for i := 0; i < 64; i++ {
		r := rand.Intn(255)
		o := uint8(r)

		a := NewPopCount(uint64(o))
		b := PopCount(uint64(o))

		t.Logf("чило %d двоичное представление %08b, число установленных битов: NewPopCount() = %d, PopCount() = %d\n", o, o, a, b)
		if a != b {
			t.Errorf("%b not equal %b\n", a, b)
		}
	}
}

func BenchmarkNewPopCount(b *testing.B) {
	r := rand.Intn(255)
	o := uint8(r)

	for i := 0; i < b.N; i++ {
		NewPopCount(uint64(o))
	}
}

func BenchmarkOldPopCount(b *testing.B) {
	r := rand.Intn(255)
	o := uint8(r)

	for i := 0; i < b.N; i++ {
		PopCount(uint64(o))
	}
}
