### Упражнение 3.4

---

Следуя подходу, использованному в примере с фигурами Лиссажу из раздела 1.7, 
создайте веб-сервер, который вычисляет поверхности и возвращает клиенту SVG-данные.
Сервер должен использовать в ответе заголовок Content-Type наподобие следующего:
```go
w.Header().Set("Content-Type", "image/svg+xml")
```
(Этот шаг не был нужен в примере с фигурами Лиссажу, так как сервер использует стандартную эвристику распознования 
распространённых форматов наподобие PNG по первым 512 байтам ответа и генерирует корректный заголовок.)
Позвольте клиенту указывать разные параметры, такие как высота, ширина и цвет, в запросе HTTP.

---

### Exercise 3.4

Following the approach of the Lissajous example in Section 1.7, construct a web server that 
computes surfaces and writes SVG data to the client. The server must set the Content-Type header like this:
```go
w.Header().Set("Content-Type", "image/svg+xml")
```

---

Флаг ```-o``` устанавливает выходную директорию бинарника:
```
$ go build -o ./bin
```
