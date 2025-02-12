// Fetchall выполняет параллельную выборку URL и сообщает о затраченном времени и размере ответа для каждого из них.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
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
		ch <- fmt.Sprintf("fetch: %v\n", err) // отправка в канал ch

		// Goexit завершает goroutine, которая его вызывает. Никакие другие goroutine не затрагиваются.
		// Goexit запускает все отложенные вызовы перед завершением goroutine.
		// Поскольку Goexit не является паникой, любые вызовы восстановления в этих отложенных функциях вернут nil.
		//Вызов Goexit из основной goroutine завершает эту goroutine без возврата func main.
		// Поскольку func main не вернулась, программа продолжает выполнение других goroutine.
		// Если все другие goroutine завершаются, программа аварийно завершается.
		runtime.Goexit() // завершаем текущую go-рутину
	}
	fmt.Printf("url: %s\n", url)

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
