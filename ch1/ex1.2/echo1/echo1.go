// Echo1 выводит аргументы командной строки

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		if i != len(os.Args)-1 {
			sep = "\n"
		} else {
			sep = ""
		}
		s += strconv.Itoa(i) + " " + os.Args[i] + sep
	}
	fmt.Println(s)
}
