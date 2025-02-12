// basename убирает компоненты каталога и суффикс типа файла.
// a => a, a.go => a, a/b/c.go => c, a/b/c.go => b.c

package basename1

func basename(s string) string {
	// отбрасываем последний символ '/' и всё перед ним.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// сохраняем всё до последней точки '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
