### Exercise 4.9:

Write a program ```wordfreq``` to report the frequency of each word in an input text file.
Call ```input.Split(bufio.ScanWords)``` before the first call to Scan to break the input into words instead of lines. 

---

### Упражнение 4.9:

Напишите программу ```wordfreq``` для подсчёта частоты каждого слова во входном текстовом файле. 
Вызовите ```input.Split(bufio.ScanWords)``` до первого вызова Scan для разбивки текста на слова, а не на строки.

---

Запускаем выполненное упражнение:

1.Переходим в нужную директорию с бинарником:
```bash
cd gopl/ch4/ex4.9/wordfreq/bin
```

2.Запускаем бинарник (используя перенаправление ввода):
```bash
$ ./wordfreq < ../cars.txt
word: Audi      freq: 1
word: WV        freq: 3
word: BMV       freq: 2
word: Fiat      freq: 1
word: Toyota    freq: 1
```

---

Вывод модульного теста:
```bash
$ go test -v
=== RUN   TestWordFreq
    wordfreq._test.go:14: map[Audi:1 BMV:2 Fiat:1 Toyota:1 WV:3]
    wordfreq._test.go:18: Audi 1
    wordfreq._test.go:18: WV 3
    wordfreq._test.go:18: BMV 2
    wordfreq._test.go:18: Fiat 1
    wordfreq._test.go:18: Toyota 1
--- PASS: TestWordFreq (0.00s)
PASS
ok      wordfreq        0.002s
```