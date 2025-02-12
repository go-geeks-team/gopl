package comma

import (
	"strings"
)

// comma Удаляет знак и десятичную точку.
// Вставляет запятые в строковое представление неотрицательного десятичного числа.
func comma(s string) string {
	v := removeSign(s) // удаляем знак (+ или -)
	v = removeDot(v)   // удаляем десятичную точку
	return commaOld(v)
}

// commaOld chapter 3.5.4 comma
func commaOld(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaOld(s[:n-3]) + "," + s[n-3:]
}

// removeDot удаляет десятичную точку
func removeDot(s string) string {
	var str string
	if strings.Contains(s, ".") {
		ss := strings.Split(s, ".")
		str = ss[0]
	} else {
		str = s
	}
	return str
}

// removeSign удаляет знаки + и -
func removeSign(s string) string {
	aft, _ := strings.CutPrefix(s, "+")
	ss, _ := strings.CutPrefix(aft, "-")
	return ss
}
