// Упражнение 2.3.
// Перепешите фкнкцию PopCount так, чтобы она использовала цикл вместо единого выражения.
// Сравните производительность двух версии.
// (В разделе 11.4 показано, как правильно сравнивать производительность различных реализаций.)

package popcount

// pc[i] — количество еденичных битов в i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения (количество установленных битов) значения x.
func PopCount(x uint64) int {
	var s int
	for i := 0; i < 8; i++ {
		s += int(pc[byte(x>>(i*8))])
	}
	return s
}

func OrigPopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}
