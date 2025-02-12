// Server 1 — минимальный "echo"- сервер

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Обработчик возвращает компонент пути из URL запроса.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}
