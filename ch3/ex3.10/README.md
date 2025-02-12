### Exercise 3.10

Write a non-recursive version of comma, using ```bytes.Buffer``` instead of string concatenation. 

---

### Упражнение 3.10

Напишите нерекурсивную версию функции comma, использующую ```bytes.Buffer``` вместо конкатенации строк. 

---

Tests output:
```bash
$ go test -v
=== RUN   TestComma1
comma_test.go:9: commaNew(122,999,888), commaOld(122,999,888)
--- PASS: TestComma1 (0.00s)
=== RUN   TestComma2
comma_test.go:19: commaNew(199,000), commaOld(199,000)
--- PASS: TestComma2 (0.00s)
=== RUN   TestComma3
comma_test.go:29: commaNew(12,000), commaOld(12,000)
--- PASS: TestComma3 (0.00s)
PASS
ok      comma_old   0.002s
```
