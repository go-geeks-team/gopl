package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var sh256 = flag.Bool("sha256", false, "")
var sh384 = flag.Bool("sha384", false, "")
var sh512 = flag.Bool("sha512", false, "")

func main() {
	flag.Parse()

	// проверка флагов командной строки
	if ok, msg := checkFlags(); !ok {
		fmt.Fprintf(os.Stderr, msg)
		os.Exit(1)
	}

	if !*sh384 && !*sh512 && *sh256 && len(os.Args) > 1 {
		for _, arg := range flag.Args() {
			fmt.Printf("argument: %s\tsha256sum: %x\n", arg, sha256.Sum256([]byte(arg)))
		}
	}
	if !*sh256 && *sh384 && !*sh512 && len(os.Args) > 1 {
		for _, arg := range flag.Args() {
			fmt.Printf("argument: %s\tsha384sum: %x\n", arg, sha512.Sum384([]byte(arg)))
		}
	}
	if !*sh256 && !*sh384 && *sh512 && len(os.Args) > 1 {
		for _, arg := range flag.Args() {
			fmt.Printf("argument: %s\tsha512sum: %x\n", arg, sha512.Sum512([]byte(arg)))
		}
	}

	if !*sh256 && !*sh384 && !*sh512 && len(os.Args) > 1 {
		for _, arg := range flag.Args() {
			fmt.Printf("argument: %s\tsha256sum: %x\n", arg, sha256.Sum256([]byte(arg)))
		}
	}
}

// checkFlags проверяет корректность введённых флагов
func checkFlags() (bool, string) {
	var flagCnt int

	if *sh256 {
		flagCnt++
	}
	if *sh384 {
		flagCnt++
	}
	if *sh512 {
		flagCnt++
	}
	if flagCnt >= 2 {
		return false, fmt.Sprint("Недопустимое количество флагов\n")
	}
	return true, ""
}
