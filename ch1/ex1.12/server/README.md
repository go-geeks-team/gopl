### Упражнение 1.12.

Измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из ```URL```.
Например, ```URL``` вида ```http://localhost:8000/?cycles=20``` устанавливает количество циклов равным 20 вместо значения по умолчанию,
равного 5. Используйте функцию ```strconv.Atoi``` для преобразования строкового параметра в целое число. 
Просмотреть документацию по данной функции можно с помощью команды ```go doc strconv.Atoi```.

Exercise 1.12:
Modify the Lissajous server to read parameter values from the URL. For example, you might arrange it so that a URL like 
```http://localhost:8000/?cycles=20``` sets the number of cycles to 20 instead of the default 5.
Use the strconv.Atoi function to convert the string parameter into an integer.
You can see its documentation with ```go doc strconv.Atoi```.

Запуск упражнения(на 8000 порте):
```
$ go run ./server.go &
```

Запуск упражнения(на другом порте):
```
$ go run ./server.go --port 7000 &
```

& — запускаем в режиме демона(отвязываемся от текущего терминала). 
