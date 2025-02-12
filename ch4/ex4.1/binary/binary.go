package binary

import (
	"crypto/sha256"
)

// diff возвращает количество различающихся битов
// возвращает 0 если аргументы идентичны
func diff(a, b []byte) int {
	var cnt int
	checksum1 := sha256.Sum256(a)
	checksum2 := sha256.Sum256(b)

	for i1, x := range checksum1 {
		for i2, y := range checksum2 {
			if i1 == i2 && x != y {
				cnt++
			}
		}
	}

	return cnt // возвращаем количество различающихся битов (возвращает 0 если аргументы идентичны)
}
