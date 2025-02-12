// Exercise 4.9.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// подсчитываем частоту каждого слова во входном текстовом файле
func wordFreq(in *os.File) map[string]int {
	ws := make(map[string]int)

	input := bufio.NewScanner(in)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		txt := input.Text()
		ws[txt]++
	}
	return ws
}

func main() {
	out := wordFreq(os.Stdin)
	for w, f := range out {
		fmt.Printf("word: %s\tfreq: %d\n", w, f)
	}
}
