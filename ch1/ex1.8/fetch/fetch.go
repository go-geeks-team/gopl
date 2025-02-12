// Упражнение 1.8
// Измените программу ```fetch``` так, чтобы к каждому аргументу ```URL``` автоматически добавлялся префикс ```http://```
// в случае отсутствия в нём такового. Можете воспользоваться функцией ```strings.HasPrefix```.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		prefix := ""

		if !strings.HasPrefix(url, "https://") {
			if !strings.HasPrefix(url, "http://") {
				prefix = "http://"
			}
		}

		resp, err := http.Get(prefix + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "чтение: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
