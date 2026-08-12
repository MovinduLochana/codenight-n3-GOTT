package main

import "testing"

func TestParsePositive(t *testing.T) {
	val, err := ParsePositive("10")
	if err != nil || val != 10 {
		t.Errorf("ParsePositive(10) failed")
	}
	_, err = ParsePositive("-5")
	if err == nil {
		t.Errorf("ParsePositive(-5) expected error")
	}
}
