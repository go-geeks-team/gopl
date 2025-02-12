### Упражнение 2.5 

Упражнение 2.5:
Выражение ```x&(x-1)``` сбрасывает крайний справа ненулевой бит ```x```.
Напишите версию PopCount, которая подсчитывает биты с использованием этого факта, 
и оцените её производительность.

Exercise 2.5
The expression ```x&(x-1)``` clears the rightmost non-zero bit of x.
Write a version of PopCount that counts bits by using this fact, and access its performance.

https://github.com/go-geeks-team/gopl/blob/9598e0fef3391fa586e7176736b3c24df2ce09fa/ch2/ex2.5/popcount/popcount.go#L1-L40

Тесты производительности:

```bash
$ go test -bench="."
goos: linux
goarch: amd64
pkg: popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkNewPopCount-12         689232790                1.725 ns/op
BenchmarkOldPopCount-12         1000000000               0.2372 ns/op
PASS
ok      popcount        1.644s
```

Модульные тесты:
```bash
$ go test -v
=== RUN   TestNewPopCount
    popcount_test.go:16: чило 254 двоичное представление 11111110, число установленных битов: NewPopCount() = 7, PopCount() = 7
    popcount_test.go:16: чило 41 двоичное представление 00101001, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 185 двоичное представление 10111001, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 4 двоичное представление 00000100, число установленных битов: NewPopCount() = 1, PopCount() = 1
    popcount_test.go:16: чило 196 двоичное представление 11000100, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 124 двоичное представление 01111100, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 65 двоичное представление 01000001, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 144 двоичное представление 10010000, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 35 двоичное представление 00100011, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 75 двоичное представление 01001011, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 20 двоичное представление 00010100, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 135 двоичное представление 10000111, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 244 двоичное представление 11110100, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 179 двоичное представление 10110011, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 6 двоичное представление 00000110, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 54 двоичное представление 00110110, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 192 двоичное представление 11000000, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 101 двоичное представление 01100101, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 55 двоичное представление 00110111, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 78 двоичное представление 01001110, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 44 двоичное представление 00101100, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 133 двоичное представление 10000101, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 10 двоичное представление 00001010, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 115 двоичное представление 01110011, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 35 двоичное представление 00100011, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 208 двоичное представление 11010000, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 251 двоичное представление 11111011, число установленных битов: NewPopCount() = 7, PopCount() = 7
    popcount_test.go:16: чило 57 двоичное представление 00111001, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 197 двоичное представление 11000101, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 247 двоичное представление 11110111, число установленных битов: NewPopCount() = 7, PopCount() = 7
    popcount_test.go:16: чило 10 двоичное представление 00001010, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 122 двоичное представление 01111010, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 173 двоичное представление 10101101, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 195 двоичное представление 11000011, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 121 двоичное представление 01111001, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 184 двоичное представление 10111000, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 181 двоичное представление 10110101, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 184 двоичное представление 10111000, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 239 двоичное представление 11101111, число установленных битов: NewPopCount() = 7, PopCount() = 7
    popcount_test.go:16: чило 158 двоичное представление 10011110, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 69 двоичное представление 01000101, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 10 двоичное представление 00001010, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 33 двоичное представление 00100001, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 118 двоичное представление 01110110, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 138 двоичное представление 10001010, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 105 двоичное представление 01101001, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 29 двоичное представление 00011101, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 229 двоичное представление 11100101, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 111 двоичное представление 01101111, число установленных битов: NewPopCount() = 6, PopCount() = 6
    popcount_test.go:16: чило 20 двоичное представление 00010100, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 42 двоичное представление 00101010, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 86 двоичное представление 01010110, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 50 двоичное представление 00110010, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 237 двоичное представление 11101101, число установленных битов: NewPopCount() = 6, PopCount() = 6
    popcount_test.go:16: чило 204 двоичное представление 11001100, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 90 двоичное представление 01011010, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 133 двоичное представление 10000101, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 137 двоичное представление 10001001, число установленных битов: NewPopCount() = 3, PopCount() = 3
    popcount_test.go:16: чило 170 двоичное представление 10101010, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 66 двоичное представление 01000010, число установленных битов: NewPopCount() = 2, PopCount() = 2
    popcount_test.go:16: чило 147 двоичное представление 10010011, число установленных битов: NewPopCount() = 4, PopCount() = 4
    popcount_test.go:16: чило 59 двоичное представление 00111011, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 110 двоичное представление 01101110, число установленных битов: NewPopCount() = 5, PopCount() = 5
    popcount_test.go:16: чило 186 двоичное представление 10111010, число установленных битов: NewPopCount() = 5, PopCount() = 5
--- PASS: TestNewPopCount (0.00s)
PASS
ok      popcount        0.002s
```