#### Exercise 1.3/Echo3

Запуск теста производительности echo3 (используем конвеер и команду tail, для вывода последних 20 строк).

```bash
$ go test -bench="." | tail -n 20 > test_out.sh
```

