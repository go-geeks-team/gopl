package strutil

import (
	"strings"
)

func TrimSpaces(s string) string {
	s = strings.TrimSpace(s)
	for i, code := range s {
		if code == 32 {
			s := s[0:i] + s[i+1:]
			return TrimSpaces(s)
		}
	}
	return s
}
