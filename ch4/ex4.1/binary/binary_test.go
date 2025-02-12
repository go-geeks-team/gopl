package binary

import (
	"testing"
)

func TestEqual(t *testing.T) {
	x1 := []byte("x")
	x2 := []byte("x")

	difCnt := diff(x1, x2) // diff возвращает количество различающихся битов
	if difCnt != 0 {
		t.Errorf("\nx1: %x\nx2:%x\ndifferent bits diff: %d\n", x1, x2, difCnt)
	}
}

func TestNotEqual(t *testing.T) {
	x1 := []byte("x")
	x2 := []byte("X")

	difCnt := diff(x1, x2) // diff возвращает количество различающихся битов
	if difCnt == 0 {
		t.Errorf("\nx1: %x\nx2:%x\ndifferent bits diff: %d\n", x1, x2, difCnt)
	}
}
