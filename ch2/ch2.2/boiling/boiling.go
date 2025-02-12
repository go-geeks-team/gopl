package main

import "fmt"

const boilingF = 212.0 // константа уровня пакета

// функция уровня пакета
func main() {
	var f = boilingF         // локальная переменная (для функции main)
	var c = (f - 32) * 5 / 9 // локальная переменная (для функции main)
	fmt.Printf("Температура кипения = %g°F или %g°C\n", f, c)
	// Запуск и вывод:
	// $ go run ./boiling.go
	//Температура кипения = 212°F или 100°C
}
