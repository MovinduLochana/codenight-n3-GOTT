package main

import "testing"

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(10)
	s.Push(20)
	if got := s.Pop(); got != 20 {
		t.Errorf("Pop = %d; want 20", got)
	}
	if got := s.Pop(); got != 10 {
		t.Errorf("Pop = %d; want 10", got)
	}
}
