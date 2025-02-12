// Echo2 выводит аргументы командной строки

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		if i < len(os.Args)-1 {
			sep = "\n"
		} else {
			sep = ""
		}
		s += strconv.Itoa(i) + " " + arg + sep
	}
	fmt.Println(s)
}
