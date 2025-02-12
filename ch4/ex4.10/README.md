### Exercise 4.10:

Modify issues to report the results in age categories, say less than a month old, less than a year old,
and more than a year old. 

---

### Упражнение 4.10:

Измените программу issues так, чтобы она сообщала о результатах с учётом их давности,
деля на категории, например, поданные менее месяца назад, менее года назад и более года назад.

---

#### Run completed exercise (Запускаем выполненное упражнение)

Compile source files to execute binary file (in folder "bin"):
```bash
$ go build  -o ./bin
```
Компилируем файлы с исходным кодом в исполняемый бинарный файл (в папку "bin")

---

#### Flags hint:

```bash
-start_page=1           — start scanning from the first page
-end_page=10            — finish scanning on page 10 (inclusive)
-period=less_month_ago  — display results submitted "less than a month ago"
-query="memory leak"    — search for issues containing the string "memory leak" in the title

$ ./issues -start_page=1 -end_page=10 -period=less_month_ago -query="memory leak"
```


#### Подсказка по флагам:

```
-start_page=1          — начинаем сканирование с первой страницы
-end_page=10           — заканчиваем сканирование на 10 странице (включительно)
-period=less_month_ago — выводим результаты поданные месяц назад ( возможные значения: "less_month_ago", "less_year_ago", "more_year_ago" )
-query=ok              — поиск issues содержащим в title строку "memory leak"
```

```bash
$ ./issues -start_page=1 -end_page=10 -period=less_month_ago -query="memory leak"
```

---

#### Run Example (Запускаем пример):


```bash
$ ./issues -period=less_month_ago -query="memory leak"
start scanning...
# Search Issues Submitted a less month ago
# Search word(s): memory leak
scanning page 1         found 1 issues
scanning page 2         found 4 issues
scanning page 3         found 4 issues
scanning page 4         found 7 issues
scanning page 5         found 4 issues
scanning page 6         found 7 issues
scanning page 7         found 4 issues
scanning page 8         found 4 issues
scanning page 9         found 8 issues
scanning page 10        found 6 issues
total issues count: 49

----------------------------------------------------------------------------------------------------
| Count |  Days ago |      Date |     Login | Title 
----------------------------------------------------------------------------------------------------
| 1    | 000000003 | 2024-12-2 | sebastien | Fix memory leak 
| 2    | 000000001 | 2024-12-3 | TheMrShel | Memory Leaks 
| 3    | 000000334 | 2024-12-3 |     solos | memory leak 
| 4    | 000000518 | 2024-12-3 | Horusiath | Memory leaks 
| 5    | 000000003 | 2024-12-1 | Nuzhny007 | Fix memory leak 
| 6    | 000000007 | 2024-12-2 |   spereku | Memory leaks 
| 7    | 000035369 | 2024-12-2 | Meteoeoeo | Memory consumption, memory leaks 
| 8    | 000005254 | 2024-12-3 |  linsaftw | Memory Leak Found 
| 9    | 000000003 | 2024-12-2 | UnityDev0 | Memory leak 
| 10   | 000000114 | 2024-12-2 |      sjjh | memory leak 
| 11   | 000000168 | 2024-12-2 |   kseyhan | Memory leak 
| 12   | 000001173 | 2024-12-3 |  idressos | ZUGFeRD memory leak 
| 13   | 000000003 | 2024-12-1 | BorisBugo | Memory leak 
| 14   | 000000002 | 2024-12-2 | ytachioka | memory leak 
| 15   | 000008276 | 2024-12-3 | selimsand | Fix metal backend memory leaks 
| 16   | 000000686 | 2024-12-1 | lievenhey | Fix memory leaks 
| 17   | 000001309 | 2024-12-2 |  kevincys | about memory leak 
| 18   | 000000016 | 2024-12-1 |       lhk | Memory Leak? 
| 19   | 000001101 | 2024-12-2 |   UnitieG | WrappedBlockState Memory Leaks 
| 20   | 000000134 | 2024-12-2 | pengtiana | Memory leak log 
| 21   | 000000153 | 2024-12-1 |   chyzwar | Memory leak 
| 22   | 000008362 | 2024-12-2 |   wangyum | Memory leak issue 
| 23   | 000000076 | 2024-12-3 | wang-jiao | [34] - memory-leaks-and-usage 
| 24   | 000000020 | 2024-12-1 | NewSoupVi | Memory Leak? 
| 25   | 000001045 | 2024-12-2 | RomanSmol | Possible memory leak 
| 26   | 000000506 | 2024-12-2 |   stav242 | Memory leak issue ? 
| 27   | 000001813 | 2024-12-2 | syamajala | Legion: memory leak 
| 28   | 000000002 | 2024-12-2 |    Lisias | Possible Memory Leak? 
| 29   | 000000015 | 2024-12-2 | Protected | GPU Memory Leak 
| 30   | 000000017 | 2024-12-1 |  quietmid | Memory Leaks 
| 31   | 000009074 | 2024-12-1 |     sp132 | [Bug]: Memory Leak 
| 32   | 000000017 | 2024-12-1 | madmiraal | Memory Leaks 
| 33   | 000000011 | 2024-12-2 |    ntlhui | Possible Memory Leak 
| 34   | 000000491 | 2024-12-1 | adilansar | [BUG] Memory leak 
| 35   | 000000037 | 2024-12-1 | yotsubo42 | [Survay] memory leak 
| 36   | 000000003 | 2024-12-0 |      Korb | Memory leak 
| 37   | 000002950 | 2024-12-3 |      CKY- | [Bug] evalJS memory leak  
| 38   | 000000916 | 2024-12-2 |   duapple | Fix memory leak in example 
| 39   | 000004208 | 2024-12-2 | xinzhongy | Node memory leak 
| 40   | 000002490 | 2024-12-2 |    abaker | Memory leak in WearCounterDataService 
| 41   | 000006988 | 2024-12-2 | matusdrob | Memory Leak Browser 
| 42   | 000007196 | 2024-12-1 | GreaterTh | Major memory leak 
| 43   | 000000001 | 2024-12-1 |    KrMzyc | CUDA memory leaking? 
| 44   | 000000001 | 2024-12-0 | HiroIshid | Memory leak 
| 45   | 000000167 | 2024-12-2 | chronicad | Detached node memory leak 
| 46   | 000000001 | 2024-12-1 |      Luzo | Potential memory leaks 
| 47   | 000000035 | 2024-12-2 |   Gudwlin | Valgrind warnings, memory leaks 
| 48   | 000003221 | 2024-12-1 | tvanderst | editor memory leak 
| 49   | 000000003 | 2024-12-2 |   ignatov | Memory leak with git 
----------------------------------------------------------------------------------------------------
start page: 1
end page: 10
per page: 10
total issues count: 49
```

---

#### Module tests output (вывод модульных тестов):
```bash
$ go test -v
=== RUN   TestLessMonthAgo
--- PASS: TestLessMonthAgo (0.00s)
=== RUN   TestLessYearAgo
--- PASS: TestLessYearAgo (0.00s)
=== RUN   TestMoreYearAgo
--- PASS: TestMoreYearAgo (0.00s)
PASS
ok      issues  0.003s
```
