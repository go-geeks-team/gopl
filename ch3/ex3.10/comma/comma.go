// Упражнение 3.10
// Напишите нерекурсивную версию функции comma_old, использующую ```bytes.Buffer``` вместо конкатенации строк.

package comma

import (
	"bytes"
)

func commaNew(s string) string {
	var buf2 bytes.Buffer

	rev := reverse([]byte(s))

	for i, x := range rev {
		if i != 0 && i%3 == 0 {
			buf2.WriteRune(',')
		}
		// fmt.Printf("%c\n", x)
		buf2.WriteRune(x)
	}

	return reverse(buf2.Bytes())
}

func commaOld(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaOld(s[:n-3]) + "," + s[n-3:]
}

func reverse(s []byte) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	d := string(s)
	return d
}
