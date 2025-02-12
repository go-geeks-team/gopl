package main

import (
	"testing"
	"time"
)

// submitted less a month ago
// отправленные менее месяца назад
func TestLessMonthAgo(t *testing.T) {
	lessMonthAgo := time.Date(time.Now().Year(), time.Now().Month()-1, time.Now().Day(), time.Now().Hour(), time.Now().Hour(), time.Now().Second(), 0, time.UTC)
	ok := submittedLessThanMonthAgo(lessMonthAgo)

	if ok {
		t.Errorf("%s\n", lessMonthAgo)
	}
}

// submitted less a year ago
// отправленные менее года назад
func TestLessYearAgo(t *testing.T) {
	lessYearAgo := time.Date(time.Now().Year(), time.Now().Month()-10, time.Now().Day(), time.Now().Hour(), time.Now().Hour(), time.Now().Second(), 0, time.UTC)
	ok := submittedLessThanOneYearAgo(lessYearAgo)

	if !ok {
		t.Errorf("%s\n", lessYearAgo)
	}
}

// submitted more a year ago
// отправленные более года назад
func TestMoreYearAgo(t *testing.T) {
	moreYearAgo := time.Date(time.Now().Year()-1, time.Now().Month()-1, time.Now().Day(), time.Now().Hour(), time.Now().Hour(), time.Now().Second(), 0, time.UTC)
	ok := submittedMoreThanYearAgo(moreYearAgo)

	if !ok {
		t.Errorf("%s\n", moreYearAgo)
	}
}
