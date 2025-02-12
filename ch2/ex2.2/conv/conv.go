// Упражнение 2.2
//
// Напишите программу общего назначения для преобразования едениц, аналогичную cf,
// которая считывает числа из аргументов командной строки (или из стандартного ввода, если аргумент командной строки отсутствует)
// и преобразует каждое число в другие еденицы, как температуру — в градусы Цельсия и Фаренгейта,
// длину — в футы и метры, вес — в фунты и килограммы и т.д.

package main

import (
	"bufio"
	"bytes"
	"conv/pressconv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var pa = flag.String("pa", "", "Converting from Pascal to Bar and Atm")
var bar = flag.String("bar", "", "Converting from Bar to Pa and Atm")
var atm = flag.String("atm", "", "Converting from Atm to Pa and Bar")

func main() {
	if len(os.Args) > 1 {
		flagHandler()

		if *pa == "" && *bar == "" && *atm == "" {
			for _, arg := range os.Args[1:] {
				conv(arg)
			}
		}
	} else {
		fmt.Print("Pressure converter: ")

		sc := bufio.NewScanner(os.Stdin)

		for sc.Scan() {
			b := sc.Bytes()
			if bytes.Equal(b, []byte("exit")) {
				break
			}

			conv(string(b))
			os.Exit(0)
		}
	}

}

func flagHandler() {
	flag.Parse()
	if *pa != "" {
		v, err := strconv.ParseFloat(*pa, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pa: %v\n", err)
			os.Exit(1)
		}

		atm := pressconv.PascalToAtm(pressconv.Pascal(v))
		b := pressconv.PascalToBar(pressconv.Pascal(v))

		fmt.Printf("%s equal %s\n", pressconv.Pascal(v), atm)
		fmt.Printf("%s equal %s\n", pressconv.Pascal(v), b)
	}

	if *bar != "" {
		v, err := strconv.ParseFloat(*bar, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bar: %v\n", err)
			os.Exit(1)
		}
		atm := pressconv.BarToAtm(pressconv.Bar(v))
		pas := pressconv.BarToPascal(pressconv.Bar(v))

		fmt.Printf("%s equal %s\n", pressconv.Bar(v), atm)
		fmt.Printf("%s equal %s\n", pressconv.Bar(v), pas)
	}

	if *atm != "" {
		v, err := strconv.ParseFloat(*atm, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "atm: %v\n", err)
			os.Exit(1)
		}
		pa := pressconv.AtmToPascal(pressconv.Atmosphere(v))
		bar := pressconv.AtmToBar(pressconv.Atmosphere(v))

		fmt.Printf("%s equal %s\n", pressconv.Atmosphere(v), pa)
		fmt.Printf("%s equal %s\n", pressconv.Atmosphere(v), bar)
	}
}
func conv(arg string) {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "atm: %v\n", err)
		os.Exit(1)
	}
	atmToPascal := pressconv.AtmToPascal(pressconv.Atmosphere(v))
	atmToBar := pressconv.AtmToBar(pressconv.Atmosphere(v))

	barToATM := pressconv.BarToAtm(pressconv.Bar(v))
	barToPascal := pressconv.BarToPascal(pressconv.Bar(v))

	pascalToAtm := pressconv.PascalToAtm(pressconv.Pascal(v))
	pascalToBar := pressconv.PascalToBar(pressconv.Pascal(v))

	fmt.Printf("%s equal %s\n", pressconv.Atmosphere(v), atmToPascal)
	fmt.Printf("%s equal %s\n", pressconv.Atmosphere(v), atmToBar)
	fmt.Println("")

	fmt.Printf("%s equal %s\n", pressconv.Bar(v), barToATM)
	fmt.Printf("%s equal %s\n", pressconv.Bar(v), barToPascal)
	fmt.Println("")

	fmt.Printf("%s equal %s\n", pressconv.Pascal(v), pascalToAtm)
	fmt.Printf("%s equal %s\n", pressconv.Pascal(v), pascalToBar)
	fmt.Println("")
}
