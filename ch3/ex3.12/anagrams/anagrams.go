// Exercise 3.12

package anagrams

import (
	"bytes"
)

func IsAnagrams(word1, word2 string) bool {
	var buf bytes.Buffer

	a := []byte(word1)
	b := []byte(word2)

	for i := len(b) - 1; i >= 0; i-- {
		buf.WriteByte(b[i])
	}
	if bytes.Equal(buf.Bytes(), a) {
		return true
	}
	return false
}
