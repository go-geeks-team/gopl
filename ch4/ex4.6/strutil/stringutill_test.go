package strutil

import "testing"

func TestTrimSpaces1(t *testing.T) {
	out := TrimSpaces("   a b cdef")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}
func TestTrimSpaces2(t *testing.T) {
	out := TrimSpaces("   a b cdef   ")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}

func TestTrimSpaces3(t *testing.T) {
	out := TrimSpaces("a b cdef")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}

func TestTrimSpaces4(t *testing.T) {
	out := TrimSpaces("a b   c d e   f")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}

func TestTrimSpaces5(t *testing.T) {
	out := TrimSpaces("a  b   c  d  e   f ")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}

func TestTrimSpaces6(t *testing.T) {
	out := TrimSpaces("abcdef")
	t.Logf("\"%s\"\n", out)
	if out != "abcdef" {
		t.Errorf("%s\n", out)
	}
}
