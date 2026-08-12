package main

import "testing"

func TestReadAge(t *testing.T) {
	val, err := ReadAge("25")
	if err != nil || val != 25 {
		t.Errorf("ReadAge(25) failed")
	}
	_, err = ReadAge("abc")
	if err == nil {
		t.Errorf("ReadAge(abc) expected error")
	}
}
