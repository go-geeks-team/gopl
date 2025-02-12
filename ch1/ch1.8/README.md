### 1.8. Некоторые мелочи.

Инструкции управления потоком

*if*, *for*, *switch* — инструкция управления потоком

*switch* — инструкция множественного ветвления

Инструкции break и continue модифицируют поток управления.

Инструкции break заставляет передать управление следующей инструкции 

continue заставляет заставляет наиболее глубоко вложенный цикл for начать очередную итерацию.

Инструкции могут иметь метки на которые можно ссылаться

Имеется также инструкция goto(предназначена для машино-генерируемого кода а не для программистов).

Объявление type позволяет присваивать имена существующим типам.

Указатель это значение содержащие адрес переменной.

Метод представляет собой функцию, связанную с именованным типом.

Интерфейсы представляют собой абстрактные типы, которые позволяют рассматривать различные конкретные типы одинаково,
на основе имеющихся у них методов.

[Предметный указатель стандартной библиотеки](https://golang.org/pkg)

[Пакеты предоставленные сообществом](https://godoc.org)

Инструмент godoc:
```
$ go doc http.ListenAndServe
package http // import "net/http"

func ListenAndServe(addr string, handler Handler) error
    ListenAndServe listens on the TCP network address addr and then calls
    Serve with handler to handle requests on incoming connections. Accepted
    connections are configured to enable TCP keep-alives.

    The handler is typically nil, in which case DefaultServeMux is used.

    ListenAndServe always returns a non-nil error.
```