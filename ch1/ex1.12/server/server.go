// Упражнение 1.12.
// Измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из ```URL```.
// Например, ```URL``` вида ```http://localhost:8000/?cycles=20``` устанавливает количество циклов равным 20 вместо значения по умолчанию,
// равного 5. Используйте функцию ```strconv.Atoi``` для преобразования строкового параметра в целое число.
// Просмотреть документацию по данной функции можно с помощью команды ```go doc strconv.Atoi```.

package main

import (
	"flag"
	"log"
	"net/http"
	"server/lissajous"
	"strconv"
)

var port = flag.String("port", "8000", "Set custom port number")

func main() {
	flag.Parse()
	http.HandleFunc("/", handleMain)
	log.Fatal(http.ListenAndServe("localhost:"+*port, nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	cycles := 5 // default param
	if r.URL.Query().Has("cycles") {
		c, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err == nil {
			cycles = c
		}
	}
	lissajous.Gen(w, cycles)
}
