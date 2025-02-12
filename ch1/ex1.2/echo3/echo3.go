// Echo3 выводит аргументы командной строки

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var sep = "\n"

	fmt.Println(strings.Join(os.Args[0:], sep))
}
