### Упражнение 2.3.

Перепешите фкнкцию PopCount так, чтобы она использовала цикл вместо единого выражения.
Сравните производительность двух версии. 
(В разделе 11.4 показано, как правильно сравнивать производительность различных реализаций.)

Unit tests output:
```bash
$ go test -v
=== RUN   TestPopCount
popcount_test.go:10: MaxValue: 18446744073709551615
--- PASS: TestPopCount (0.00s)
PASS
ok      popcount        0.002s
```

Benchmarks output:
```bash
$ go test -v -bench="."
=== RUN   TestPopCount
popcount_test.go:10: MaxValue: 18446744073709551615
--- PASS: TestPopCount (0.00s)
goos: linux
goarch: amd64
pkg: popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkOrigPopCount
BenchmarkOrigPopCount-12        1000000000               0.2290 ns/op
BenchmarkPopCount
BenchmarkPopCount-12            166061696                7.228 ns/op
PASS
ok      popcount        2.193s
```