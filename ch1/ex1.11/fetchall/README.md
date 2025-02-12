### Упражнение 1.11

Выполните ```fetchall``` с длинным списком аргументов, таким как образцы, доступные на сайте ```alexa.com```.
Как ведёт себя программа, когда веб-сайт просто не отвечает? (В разделе 8.9 описан механизм отслеживания таких ситуаций.)

Exercise 1.11:
Try fetchall with longer argument lists, such as samples from the top million web sites
available at alexa.com. How does the program behave if a web site just doesn’t respond? (Section 8.9
describes mechanisms for coping in such cases.)

close go-routine

### 1.6. Параллельная выборка URL

```
$ go run ./fetchall.go https://golang.org https://gopl.io https://godoc.org
1.34s   62384 https://golang.org
1.40s    4154 https://gopl.io
1.40s   32284 https://godoc.org
1.40s elapseed
```
