package main

import (
	"os"
	"testing"
)

func TestWordFreq(t *testing.T) {
	f, err := os.Open("cars.txt")
	if err != nil {
		t.Errorf("os.Open: %s\n", err)
	}
	wf := wordFreq(f)
	t.Logf("%v\n", wf)

	wvCnt := 0
	for w, f := range wf {
		t.Logf("%s %d", w, f)
		if w == "WV" && f != 3 {
			t.Errorf("WV not equal 3 equal %d\n", wvCnt)
		}
	}
}
