package popcount

import (
	"testing"
)

const MaxValue = ^uint64(0) // uint64 max value

func TestPopCount(t *testing.T) {
	t.Logf("MaxValue: %d\n", MaxValue)

	if OrigPopCount(MaxValue) != PopCount(18446744073709551615) {
		t.Errorf("%v not equal %v\n", OrigPopCount(18446744073709551615), PopCount(MaxValue))
	}

	if OrigPopCount(18446744073709551615) != PopCount(18446744073709551615) {
		t.Errorf("%v should be equal %v\n", OrigPopCount(18446744073709551615), PopCount(18446744073709551615))
	}

	if OrigPopCount(1000) != PopCount(1000) {
		t.Errorf("%v should be equal %v\n", PopCount(1000), OrigPopCount(1000))
	}
}

func BenchmarkOrigPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrigPopCount(uint64(i))
	}
}
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
