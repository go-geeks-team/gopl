package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	out := comma("122999")
	t.Logf("%s\n", out)
	if out != "122,999" {
		t.Errorf("%s\n", out)
	}
}
