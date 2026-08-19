package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Pop(); got != 3 {
		t.Errorf("first Pop = %d; want 3", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("second Pop = %d; want 2", got)
	}
	if !reflect.DeepEqual(s.Items, []int{1}) {
		t.Errorf("after two pops, Items = %v; want [1]", s.Items)
	}

	s.Push(9)
	if got := s.Pop(); got != 9 {
		t.Errorf("third Pop = %d; want 9", got)
	}
}