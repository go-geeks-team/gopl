// Упражнение 1.9.
// Измените программу fetch так, чтобы она выводила код состояния ```HTTP```, содержащийся в ```resp.Status```.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "чтение: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("<!-- response status: %s -->\n", resp.Status)
		fmt.Printf("%s", b)
	}
}
