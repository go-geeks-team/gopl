package main

import (
	"testing"
	"time"
)

func BenchmarkMain(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {
		main()
	}
	sec := time.Since(start).Seconds()
	b.Logf("%.6f Sec.", sec)
}
