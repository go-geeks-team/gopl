// Exercise 4.8.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Количество символов Unicode
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8

	letterCount := make(map[rune]int) // Количество букв Unicode
	digitCount := make(map[rune]int)
	spacesCount := make(map[rune]int)
	upperCount := make(map[rune]int)
	graphicCount := make(map[rune]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // Возврат руны, байтов, ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		if unicode.IsLetter(r) {
			letterCount[r]++
		}
		if unicode.IsDigit(r) {
			digitCount[r]++
		}
		if unicode.IsSpace(r) {
			spacesCount[r]++
		}
		if unicode.IsUpper(r) {
			upperCount[r]++
		}

		if unicode.IsGraphic(r) {
			graphicCount[r]++
		}
	}
	fmt.Printf("graphic count\tcount\n")
	for c, n := range graphicCount {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("upper count\tcount\n")
	for c, n := range upperCount {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("spaces\tcount\n")
	for c, n := range spacesCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("letter\tcount\n")
	for c, n := range letterCount {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("digit\tcount\n")
	for c, n := range digitCount {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}
