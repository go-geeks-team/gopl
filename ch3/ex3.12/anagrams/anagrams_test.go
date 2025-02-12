package anagrams

import "testing"

func TestA(t *testing.T) {
	s1 := "super"
	s2 := "repus"

	ok := IsAnagrams(s1, s2)
	if !ok {
		t.Errorf("%s not anagrama of %s\n", s1, s2)
	}
}
func TestB(t *testing.T) {
	s1 := "super"
	s2 := "abcde"

	ok := IsAnagrams(s1, s2)
	if ok {
		t.Errorf("%s not anagrama of %s\n", s1, s2)
	}
}

func TestC(t *testing.T) {
	s1 := "super"
	s2 := "super"

	ok := IsAnagrams(s1, s2)
	if ok {
		t.Errorf("%s not anagrama of %s\n", s1, s2)
	}
}
func TestD(t *testing.T) {
	s1 := "repus"
	s2 := "super"

	ok := IsAnagrams(s1, s2)
	if !ok {
		t.Errorf("%s is anagram of %s\n", s1, s2)
	}
}
