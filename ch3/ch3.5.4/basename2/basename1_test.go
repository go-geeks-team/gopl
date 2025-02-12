package basename2

import (
	"testing"
)

func TestBasename1(t *testing.T) {
	b := basename("a")
	if b != "a" {
		t.Errorf("%s not equal a\n", b)
	}
}

func TestBasename2(t *testing.T) {
	b := basename("a.go")
	if b != "a" {
		t.Errorf("%s not equal a\n", b)
	}
}

func TestBasename3(t *testing.T) {
	b := basename("a/b/c.go")
	if b != "c" {
		t.Errorf("%s not equal c\n", b)
	}
}
