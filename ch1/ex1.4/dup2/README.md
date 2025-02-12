### Упражнение 1.4.

Измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.

Пример запуска dup2(используя перенаправления вывода в файл):
```bash
$ go run ./dup2.go in1.txt in2.txt in3.txt > out.txt
```

Пример запуска dup2 с выводом в stdout:
```bash
$ go run ./dup2.go in1.txt in2.txt in3.txt
2       in2.txt/Astana
3       in1.txt/bmw
2       in1.txt/subaru
4       in3.txt/Lenovo
```