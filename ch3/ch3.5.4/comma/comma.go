// comma вставляет запятые в строковое представление неотрицательного десятичного числа.

package comma

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
