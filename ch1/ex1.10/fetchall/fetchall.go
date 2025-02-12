// Fetchall выполняет параллельную выборку URL и сообщает о затраченном времени и размере ответа для каждого из них.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // запуск go-подпрограммы
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		// получение из канала ch
	}
	fmt.Printf("%.2fs elapseed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // отправка в канал ch
		return
	}
	f, err := os.Create("out.html")
	if err != nil {
		ch <- fmt.Sprintf("os.Create %s: %v", url, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
