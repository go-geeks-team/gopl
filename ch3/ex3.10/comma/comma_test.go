package comma

import "testing"

func TestComma1(t *testing.T) {
	c1 := commaNew("122999888")
	c2 := commaOld("122999888")

	t.Logf("commaNew(%v), commaOld(%v)", c1, c2)
	if c1 != c2 {
		t.Errorf("commaNew(%s) not equal commaOld(%s)\n", c1, c2)
	}
}

func TestComma2(t *testing.T) {
	c1 := commaNew("199000")
	c2 := commaOld("199000")

	t.Logf("commaNew(%v), commaOld(%v)", c1, c2)
	if c1 != c2 {
		t.Errorf("commaNew((%v) not equal commOld(%v)", c1, c2)
	}
}

func TestComma3(t *testing.T) {
	c1 := commaNew("12000")
	c2 := commaOld("12000")

	t.Logf("commaNew(%v), commaOld(%v)", c1, c2)
	if c1 != c2 {
		t.Errorf("commaNew((%v) not equal commOld(%v)", c1, c2)
	}
}
