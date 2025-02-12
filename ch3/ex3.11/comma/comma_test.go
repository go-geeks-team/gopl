package comma

import "testing"

// Test with sign and dot
func TestComma1(t *testing.T) {
	out := comma("+122000000.1")
	t.Logf("App: %s\n", out)
	if out != "122,000,000" {
		t.Errorf("TestComma1 Failed: %s\n", out)
	}
}

// Test with sign
func TestComma2(t *testing.T) {
	out := comma("+122000000")
	t.Logf("App: %s\n", out)
	if out != "122,000,000" {
		t.Errorf("TestComma2 Failed: %s\n", out)
	}
}

// Test dot
func TestComma3(t *testing.T) {
	out := comma("122000000.1")
	t.Logf("App: %s\n", out)
	if out != "122,000,000" {
		t.Errorf("TestComma3 Failed: %s\n", out)
	}
}

// Test without sign and dot
func TestComma4(t *testing.T) {
	out := comma("122000000")
	t.Logf("App: %s\n", out)
	if out != "122,000,000" {
		t.Errorf("TestComma4 Failed: %s\n", out)
	}
}
