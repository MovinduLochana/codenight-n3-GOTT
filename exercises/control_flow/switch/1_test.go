package main

import "testing"

func TestDayType(t *testing.T) {
	if got := DayType("Sat"); got != "Weekend" {
		t.Errorf("DayType(Sat) = %q; want Weekend", got)
	}
	if got := DayType("Mon"); got != "Weekday" {
		t.Errorf("DayType(Mon) = %q; want Weekday", got)
	}
}
