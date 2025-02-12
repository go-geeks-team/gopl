package sliceutil

import (
	"testing"
)

// тест проверяет удаление всех смежных дубликатов в срезе s1
func TestRemDup(t *testing.T) {
	s1 := []string{"audi", "wv", "mercedes", "bmw", "wv", "wv", "bmv", "audi", "audi"} // срез со смежными дубликатами
	s2 := []string{"audi", "wv", "mercedes", "bmw", "wv", "bmv", "audi"}               // срез без смежных дубликатов
	t.Logf("s1: %v\n", s1)
	t.Logf("s2: %v\n", s2)

	RemDup(s1)

	cnt := 0
	for i1, v1 := range s1 {
		for i2, v2 := range s2 {
			if i1 == i2 && v1 == v2 {
				cnt++
			}
		}
	}
	if cnt != len(s2) {
		t.Errorf("%d not equal %d", cnt, len(s2))
	}
}
