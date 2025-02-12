### Упражнение 2.2

Напишите программу общего назначения для преобразования едениц, аналогичную cf,
которая считывает числа из аргументов командной строки (или из стандартного ввода, если аргумент командной строки отсутствует)
и преобразует каждое число в другие еденицы, как температуру — в градусы Цельсия и Фаренгейта,
длину — в футы и метры, вес — в фунты и килограммы и т.д.

---

Usage without flags:
```bash
$ go run ./conv.go 3
3.000000000 atm equal 303900.00 Pa
3.000000000 atm equal 3.039000000 Bar

3.000000000 Bar equal 2.961500494 atm
3.000000000 Bar equal 300000.00 Pa

3.00 Pa equal 0.000029608 atm
3.00 Pa equal 0.000030000 Bar
```

Usage without flags and arguments:
```bash
$ go run ./conv.go
Pressure converter: 3
3.000000 atm equal 303900.00 Pa
3.000000 atm equal 3.039000 Bar

3.000000 Bar equal 2.961500 atm
3.000000 Bar equal 300000.00 Pa

3.00 Pa equal 0.000030 atm
3.00 Pa equal 0.000030 Bar
```

---

Usage with flags:
```bash
$ go run ./conv.go -bar 3
3.000000000 Bar equal 2.961500494 atm
3.000000000 Bar equal 300000.00 Pa
```

```bash
$ go run ./conv.go -atm 3
3.000000000 atm equal 303900.00 Pa
3.000000000 atm equal 3.039000000 Bar
```

```bash
$ go run ./conv.go -pa 3
3.00 Pa equal 0.000029608 atm
3.00 Pa equal 0.000030000 Bar
```

