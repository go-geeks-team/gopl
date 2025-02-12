#### Exercise 1.3/Echo1

Запуск теста производительности echo1 (используем конвеер и команду tail, для вывода последних 20 строк).

```bash
$ go test -bench="." | tail -n 20 > test_out.sh
```