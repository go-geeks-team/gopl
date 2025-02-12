package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"srv/mandelbrot"
)

var port = flag.String("port", "8000", "Default port is 8000")

func main() {
	flag.Parse()
	fmt.Fprintf(os.Stdout, "Сервер запущен на порту: %s\n", *port)
	http.HandleFunc("/", mainHandler)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mandelbrot.Gen(w, r)
}
