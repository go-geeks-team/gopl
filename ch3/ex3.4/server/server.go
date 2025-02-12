package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "8000", "флаг устанавливает http-порт (по умолчанию 8000)")

func main() {
	flag.Parse()
	fmt.Fprintf(os.Stdout, "Server started on the port %s\n", *port)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml") // устанавливаем неоюходимый HTTP-заголовок
	Gen(w)
}
